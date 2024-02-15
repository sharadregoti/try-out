// import Pipeline from "./pipeline.mjs";

// const p = new Pipeline();
// p.java()
// p.run()

import JavaPipeline from "./java.mjs";
const p = new JavaPipeline("app");
p.installDependencies().gradle(":8.3.0-jdk17-alpine")
p.build().gradle(":8.3.0-jdk17-alpine")
p.run()