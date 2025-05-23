const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Start a Container Apps Job
 *
 * @summary Start a Container Apps Job
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2025-01-01/examples/Job_Start.json
 */
async function runAContainerAppsJob() {
  const subscriptionId =
    process.env["APPCONTAINERS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["APPCONTAINERS_RESOURCE_GROUP"] || "rg";
  const jobName = "testcontainerappsjob0";
  const template = {
    containers: [
      {
        name: "testcontainerappsjob0",
        image: "repo/testcontainerappsjob0:v4",
        resources: { cpu: 0.5, memory: "1Gi" },
      },
    ],
    initContainers: [
      {
        name: "testinitcontainerAppsJob0",
        args: ["-c", "while true; do echo hello; sleep 10;done"],
        command: ["/bin/sh"],
        image: "repo/testcontainerappsjob0:v4",
        resources: { cpu: 0.5, memory: "1Gi" },
      },
    ],
  };
  const options = { template };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.jobs.beginStartAndWait(resourceGroupName, jobName, options);
  console.log(result);
}
