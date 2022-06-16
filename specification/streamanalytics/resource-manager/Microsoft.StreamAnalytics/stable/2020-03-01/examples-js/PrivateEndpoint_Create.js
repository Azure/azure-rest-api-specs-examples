const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Stream Analytics Private Endpoint or replaces an already existing Private Endpoint.
 *
 * @summary Creates a Stream Analytics Private Endpoint or replaces an already existing Private Endpoint.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/PrivateEndpoint_Create.json
 */
async function createAPrivateEndpoint() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "sjrg";
  const clusterName = "testcluster";
  const privateEndpointName = "testpe";
  const privateEndpoint = {
    manualPrivateLinkServiceConnections: [
      {
        groupIds: ["groupIdFromResource"],
        privateLinkServiceId:
          "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateLinkServices/testPls",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.privateEndpoints.createOrUpdate(
    resourceGroupName,
    clusterName,
    privateEndpointName,
    privateEndpoint
  );
  console.log(result);
}

createAPrivateEndpoint().catch(console.error);
