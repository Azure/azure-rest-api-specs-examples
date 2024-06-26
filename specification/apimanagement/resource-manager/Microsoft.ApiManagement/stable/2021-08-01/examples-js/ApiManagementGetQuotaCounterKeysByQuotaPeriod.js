const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the value of the quota counter associated with the counter-key in the policy for the specific period in service instance.
 *
 * @summary Gets the value of the quota counter associated with the counter-key in the policy for the specific period in service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetQuotaCounterKeysByQuotaPeriod.json
 */
async function apiManagementGetQuotaCounterKeysByQuotaPeriod() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const quotaCounterKey = "ba";
  const quotaPeriodKey = "0_P3Y6M4DT12H30M5S";
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.quotaByPeriodKeys.get(
    resourceGroupName,
    serviceName,
    quotaCounterKey,
    quotaPeriodKey
  );
  console.log(result);
}
