const { EventHubManagementClient } = require("@azure/arm-eventhub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a namespace. Once created, this namespace's resource manifest is immutable. This operation is idempotent.
 *
 * @summary creates or updates a namespace. Once created, this namespace's resource manifest is immutable. This operation is idempotent.
 * x-ms-original-file: 2026-01-01/NameSpaces/NamespaceWithGeoDRCreate.json
 */
async function namespaceWithGeoDRCreate() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "SampleSubscription";
  const client = new EventHubManagementClient(credential, subscriptionId);
  const result = await client.namespaces.createOrUpdate(
    "ResurceGroupSample",
    "NamespaceGeoDRCreateSample",
    {
      location: "East US",
      geoDataReplication: {
        locations: [
          { locationName: "eastus", roleType: "Primary" },
          { locationName: "westus", roleType: "Secondary" },
          { locationName: "centralus", roleType: "Secondary" },
        ],
        maxReplicationLagDurationInSeconds: 60,
      },
    },
  );
  console.log(result);
}
