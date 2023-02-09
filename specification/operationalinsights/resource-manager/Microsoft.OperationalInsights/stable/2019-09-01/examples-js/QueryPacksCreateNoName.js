const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Log Analytics QueryPack. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 *
 * @summary Creates a Log Analytics QueryPack. Note: You cannot specify a different value for InstrumentationKey nor AppId in the Put operation.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2019-09-01/examples/QueryPacksCreateNoName.json
 */
async function queryPackCreateNoName() {
  const subscriptionId =
    process.env["OPERATIONALINSIGHTS_SUBSCRIPTION_ID"] || "86dc51d3-92ed-4d7e-947a-775ea79b4919";
  const resourceGroupName =
    process.env["OPERATIONALINSIGHTS_RESOURCE_GROUP"] || "my-resource-group";
  const logAnalyticsQueryPackPayload = {
    location: "South Central US",
  };
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.queryPacks.createOrUpdateWithoutName(
    resourceGroupName,
    logAnalyticsQueryPackPayload
  );
  console.log(result);
}
