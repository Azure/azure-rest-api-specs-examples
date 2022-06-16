const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches a Managed Environment using JSON Merge Patch
 *
 * @summary Patches a Managed Environment using JSON Merge Patch
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ManagedEnvironments_Patch.json
 */
async function patchManagedEnvironment() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "examplerg";
  const environmentName = "testcontainerenv";
  const environmentEnvelope = {
    location: "East US",
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironments.beginUpdateAndWait(
    resourceGroupName,
    environmentName,
    environmentEnvelope
  );
  console.log(result);
}

patchManagedEnvironment().catch(console.error);
