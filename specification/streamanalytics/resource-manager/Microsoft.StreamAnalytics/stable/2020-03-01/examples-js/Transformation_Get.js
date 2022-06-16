const { StreamAnalyticsManagementClient } = require("@azure/arm-streamanalytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details about the specified transformation.
 *
 * @summary Gets details about the specified transformation.
 * x-ms-original-file: specification/streamanalytics/resource-manager/Microsoft.StreamAnalytics/stable/2020-03-01/examples/Transformation_Get.json
 */
async function getATransformation() {
  const subscriptionId = "56b5e0a9-b645-407d-99b0-c64f86013e3d";
  const resourceGroupName = "sjrg4423";
  const jobName = "sj8374";
  const transformationName = "transformation952";
  const credential = new DefaultAzureCredential();
  const client = new StreamAnalyticsManagementClient(credential, subscriptionId);
  const result = await client.transformations.get(resourceGroupName, jobName, transformationName);
  console.log(result);
}

getATransformation().catch(console.error);
