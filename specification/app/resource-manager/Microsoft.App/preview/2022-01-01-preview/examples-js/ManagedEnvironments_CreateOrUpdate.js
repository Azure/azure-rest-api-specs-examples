const { ContainerAppsAPIClient } = require("@azure/arm-app");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Managed Environment used to host container apps.
 *
 * @summary Creates or updates a Managed Environment used to host container apps.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/preview/2022-01-01-preview/examples/ManagedEnvironments_CreateOrUpdate.json
 */
async function createEnvironments() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "examplerg";
  const name = "testcontainerenv";
  const environmentEnvelope = {
    appLogsConfiguration: {
      logAnalyticsConfiguration: { customerId: "string", sharedKey: "string" },
    },
    location: "East US",
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    name,
    environmentEnvelope
  );
  console.log(result);
}

createEnvironments().catch(console.error);
