const { ContainerAppsAPIClient } = require("@azure/arm-app");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Patches a Managed Environment. Only patching of tags is supported currently
 *
 * @summary Patches a Managed Environment. Only patching of tags is supported currently
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/ManagedEnvironments_Patch.json
 */
async function patchManagedEnvironment() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "examplerg";
  const name = "testcontainerenv";
  const environmentEnvelope = {
    tags: { tag1: "value1", tag2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironments.update(
    resourceGroupName,
    name,
    environmentEnvelope
  );
  console.log(result);
}

patchManagedEnvironment().catch(console.error);
