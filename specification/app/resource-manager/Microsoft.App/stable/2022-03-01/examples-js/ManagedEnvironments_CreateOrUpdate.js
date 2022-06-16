const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a Managed Environment used to host container apps.
 *
 * @summary Creates or updates a Managed Environment used to host container apps.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ManagedEnvironments_CreateOrUpdate.json
 */
async function createEnvironments() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "examplerg";
  const environmentName = "testcontainerenv";
  const environmentEnvelope = {
    appLogsConfiguration: {
      logAnalyticsConfiguration: { customerId: "string", sharedKey: "string" },
    },
    daprAIConnectionString:
      "InstrumentationKey=00000000-0000-0000-0000-000000000000;IngestionEndpoint=https://northcentralus-0.in.applicationinsights.azure.com/",
    location: "East US",
    zoneRedundant: true,
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironments.beginCreateOrUpdateAndWait(
    resourceGroupName,
    environmentName,
    environmentEnvelope
  );
  console.log(result);
}

createEnvironments().catch(console.error);
