import { connect, Client, Container, Directory } from "@dagger.io/dagger"

// This uses TypeScript syntax for interfaces
interface Pipeline {
    installDependencies(
        applicationRootDirectory?: string
    ): void;
    build(): void;
    lint(): void;
    test(): void;
    sca(): void;
    run(): void;
}

class AutoPipelineContainer {
    private from: string;
    private command: string;

    constructor(from?: string, command?: string) {
        this.from = from || ""
        this.command = command || ""
    }

    getFrom(): string {
        return this.from
    }

    getCommand(): string {
        return this.command
    }
}

class BuildDaggerContainer {
    private c: Client
    private apc: AutoPipelineContainer;
    private applicationRootDirectory: string;

    constructor(rc: AutoPipelineContainer, c: Client, applicationRootDirectory?: string) {
        this.apc = rc
        this.c = c
        this.applicationRootDirectory = applicationRootDirectory || "";

    }

    runContainer(): Container {
        const source = this.c.host().directory(".", {})

        let contd: Container = this.c.container()
        contd.from(this.apc.getFrom())
            .withDirectory("/mnt", source)
            .withWorkdir("/mnt/" + this.applicationRootDirectory)
            .withEntrypoint(["bash", "-c"])
            .withExec([this.apc.getCommand()])
            .stdout()
        return contd
    }
}

class JavaPipeline implements Pipeline {
    // @ts-ignore
    private idb: InstallDependenciesBuilder;
    // @ts-ignore
    private bb: BuildBuilder;

    private applicationRootDirectory: string;
    constructor(applicationRootDirectory?: string) {
        this.applicationRootDirectory = applicationRootDirectory || "";
    }

    installDependencies(): InstallDependenciesBuilder {
        this.idb = new InstallDependenciesBuilder();
        return this.idb;
    }

    build(): BuildBuilder {
        this.bb = new BuildBuilder();
        return this.bb;
    }

    lint() {
        // Java-specific linting logic
    }

    test() {
        // Java-specific linting logic
    }

    sca() {
        // Java-specific static code analysis logic
    }

    run() {
        connect(
            async (client: Client) => {
                // get reference to the local project

                const bc: BuildDaggerContainer = new BuildDaggerContainer(this.idb.getContainer(), client, this.applicationRootDirectory)
                bc.runContainer()

                const bbc: BuildDaggerContainer = new BuildDaggerContainer(this.bb.getContainer(), client, this.applicationRootDirectory)
                bbc.runContainer()

            },
            { LogOutput: process.stderr }
        )
    }
}


class InstallDependenciesBuilder {
    // @ts-ignore
    private container: AutoPipelineContainer;

    constructor() {
    }

    maven(): void {
        // Maven-specific build logic
    }

    gradle(
        version?: string,
        command?: string
    ): void {

        command = command || "gradle dependencies"
        version = version || ""
        const from: string = 'gradle' + version
        this.container = new AutoPipelineContainer(from, command)
    }

    getContainer(): AutoPipelineContainer {
        return this.container
    }
}


class BuildBuilder {
    // @ts-ignore
    private container: AutoPipelineContainer;

    constructor() {
    }

    maven(): void {
        // Maven-specific build logic
    }

    gradle(
        version?: string,
        command?: string
    ): void {

        command = command || "gradle clean build"
        version = version || ""
        const from: string = 'gradle' + version
        this.container = new AutoPipelineContainer(from, command)
    }

    getContainer(): AutoPipelineContainer {
        return this.container
    }
}

export default JavaPipeline