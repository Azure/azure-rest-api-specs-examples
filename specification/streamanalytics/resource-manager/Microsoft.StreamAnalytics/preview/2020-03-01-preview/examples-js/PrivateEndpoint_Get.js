const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets information about the specified Private Endpoint.
 *
 * @summary Gets information about the specified Private Endpoint.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/preview/2020-03-01-preview/examples/PrivateEndpoint_Get.json
 */
async function getAPrivateEndpoint() {
  const subscriptionId =
    process.env["STREAMANALYTICS_SUBSCRIPTION_ID"] || "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = process.env["STREAMANALYTICS_RESOURCE_GROUP"] || "sjrg";
  const clusterName = "testcluster";
  const privateEndpointName = "testpe";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.privateEndpoints.get(
    resourceGroupName,
    clusterName,
    privateEndpointName,
  );
  console.log(result);
}
