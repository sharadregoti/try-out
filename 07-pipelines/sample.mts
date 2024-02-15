import { connect, Client } from "@dagger.io/dagger"

// initialize Dagger client
connect(
    async (client: Client) => {
        // get reference to the local project
        const source = client.host().directory(".", {})

        // get gradle image
        const node = client.container().from("openjdk")

        // const script = `ls; cd services/notification-service; ls;`

        const script = `ls; cd services/notification-service; echo "Staring build"; ./gradlew clean build --no-watch-fs;`

        // mount cloned repository into Node image
        const runner = node
            .withDirectory("/mnt", source)
            .withWorkdir("/mnt")
            .withExec(["bash", "-c", script])
            .stdout()

        const sca = client.container().from("aquasec/trivy")

        const sca_script = `cd services/notification-service; echo "Started scanning...."; trivy fs -s="CRITICAL,HIGH" .;`

        const sca_job = sca
            .withDirectory("/mnt", source)
            .withWorkdir("/mnt")
            .withEntrypoint(["sh"])
            .withExec(["-c", sca_script])
            .stdout()

    },
    { LogOutput: process.stderr }
)