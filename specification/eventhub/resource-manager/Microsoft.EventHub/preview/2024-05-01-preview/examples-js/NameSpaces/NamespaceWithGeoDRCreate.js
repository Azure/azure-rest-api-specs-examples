const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates a namespace. Once created, this namespace's resource manifest is immutable. This operation is idempotent.
 *
 * @summary Creates or updates a namespace. Once created, this namespace's resource manifest is immutable. This operation is idempotent.
 * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/preview/2024-05-01-preview/examples/NameSpaces/NamespaceWithGeoDRCreate.json
 */
async function namespaceWithGeoDrCreate() {
  const subscriptionId = process.env["EVENTHUB_SUBSCRIPTION_ID"] || "SampleSubscription";
  const resourceGroupName = process.env["EVENTHUB_RESOURCE_GROUP"] || "ResurceGroupSample";
  const namespaceName = "NamespaceGeoDRCreateSample";
  const parameters = {
    geoDataReplication: {
      locations: [
        { locationName: "eastus", roleType: "Primary" },
        { locationName: "westus", roleType: "Secondary" },
        { locationName: "centralus", roleType: "Secondary" },
      ],
      maxReplicationLagDurationInSeconds: 60,
    },
    location: "East US",
  };
  const credential = new DefaultAzureCredential();
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.namespaces.beginCreateOrUpdateAndWait(
    resourceGroupName,
    namespaceName,
    parameters,
  );
  console.log(result);
}
