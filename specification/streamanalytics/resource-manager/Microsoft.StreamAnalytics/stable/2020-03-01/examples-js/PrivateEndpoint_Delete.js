const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the specified private endpoint.
 *
 * @summary Delete the specified private endpoint.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/PrivateEndpoint_Delete.json
 */
async function deleteAPrivateEndpoint() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "sjrg";
  const clusterName = "testcluster";
  const privateEndpointName = "testpe";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.privateEndpoints.beginDeleteAndWait(
    resourceGroupName,
    clusterName,
    privateEndpointName
  );
  console.log(result);
}

deleteAPrivateEndpoint().catch(console.error);
