import { connect, Client, Container, Directory } from "@dagger.io/dagger"

class Build {
    public base_image: string;
    public script: string;

    constructor(base_image?: string, script?: string) {
        this.base_image = base_image || "";
        this.script = script || "";
    }
}

class Pipeline {
    private java_build_container: Build;

    constructor() {
        this.java_build_container = new Build()
    }

    // Method for Java
    java(): void {
        let c: Container = new Container;

        // set gradle image
        const node = c.from("openjdk")

        // const script = `ls; cd services/notification-service; ls;`

        const script = `ls; cd services/notification-service; echo "Staring build"; ./gradlew clean build --no-watch-fs;`

        // mount cloned repository into Node image
        // this.java_container = node.withExec(["bash", "-c", script])
        this.java_build_container = new Build("openjdk", script)
    }

    // Method for Golang
    golang(): void {
        console.log("Golang method invoked!");
    }



    run(): void {
        connect(
            async (client: Client) => {
                // get reference to the local project
                const source = client.host().directory(".", {})

                client.container().with(buildContainer(this.java_build_container, source))
            },
            { LogOutput: process.stderr }
        )
    }
}

function buildContainer(b: Build, source: Directory) {
    return (c: Container): Container => {

        c.from(b.base_image)
            .withExec(["bash", "-c", b.script])
            .withDirectory("/mnt", source)
            .withWorkdir("/mnt")
            .stdout()
        return c
    }
}

export default Pipeline