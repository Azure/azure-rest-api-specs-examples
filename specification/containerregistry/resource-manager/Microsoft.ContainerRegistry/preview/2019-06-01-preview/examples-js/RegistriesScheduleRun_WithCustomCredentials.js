const { ContainerRegistryManagementClient } = require("@azure/arm-containerregistry");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Schedules a new run based on the request parameters and add it to the run queue.
 *
 * @summary Schedules a new run based on the request parameters and add it to the run queue.
 * x-ms-original-file: specification/containerregistry/resource-manager/Microsoft.ContainerRegistry/preview/2019-06-01-preview/examples/RegistriesScheduleRun_WithCustomCredentials.json
 */
async function registriesScheduleRunWithCustomCredentials() {
  const subscriptionId =
    process.env["CONTAINERREGISTRY_SUBSCRIPTION_ID"] || "4385cf00-2d3a-425a-832f-f4285b1c9dce";
  const resourceGroupName = process.env["CONTAINERREGISTRY_RESOURCE_GROUP"] || "myResourceGroup";
  const registryName = "myRegistry";
  const runRequest = {
    type: "DockerBuildRequest",
    agentConfiguration: { cpu: 2 },
    arguments: [
      { name: "mytestargument", isSecret: false, value: "mytestvalue" },
      {
        name: "mysecrettestargument",
        isSecret: true,
        value: "mysecrettestvalue",
      },
    ],
    credentials: {
      customRegistries: {
        myregistryAzurecrIo: {
          password: { type: "Opaque", value: "***" },
          userName: { type: "Opaque", value: "reg1" },
        },
        myregistry2AzurecrIo: {
          password: { type: "Opaque", value: "***" },
          userName: { type: "Opaque", value: "reg2" },
        },
      },
      sourceRegistry: { loginMode: "Default" },
    },
    dockerFilePath: "DockerFile",
    imageNames: ["azurerest:testtag"],
    isArchiveEnabled: true,
    isPushEnabled: true,
    noCache: true,
    platform: { architecture: "amd64", os: "Linux" },
    sourceLocation:
      "https://myaccount.blob.core.windows.net/sascontainer/source.zip?sv=2015-04-05&st=2015-04-29T22%3A18%3A26Z&se=2015-04-30T02%3A23%3A26Z&sr=b&sp=rw&sip=168.1.5.60-168.1.5.70&spr=https&sig=Z%2FRHIX5Xcg0Mq2rqI3OlWTjEg2tYkboXr1P9ZUXDtkk%3D",
    target: "stage1",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerRegistryManagementClient(credential, subscriptionId);
  const result = await client.registries.beginScheduleRunAndWait(
    resourceGroupName,
    registryName,
    runRequest
  );
  console.log(result);
}
