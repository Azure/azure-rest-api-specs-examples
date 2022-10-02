const { ContainerAppsAPIClient } = require("@azure/arm-app");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches a Container App. Currently only patching of tags is supported
 *
 * @summary Patches a Container App. Currently only patching of tags is supported
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/ContainerApps_Patch.json
 */
async function patchContainerApp() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "rg";
  const name = "testcontainerApp0";
  const containerAppEnvelope = {
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.containerApps.update(resourceGroupName, name, containerAppEnvelope);
  console.log(result);
}

patchContainerApp().catch(console.error);
