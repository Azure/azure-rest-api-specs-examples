const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Updates an existing quota counter value in the specified service instance.
 *
 * @summary Updates an existing quota counter value in the specified service instance.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementUpdateQuotaCounterKeyByQuotaPeriod.json
 */
async function apiManagementUpdateQuotaCounterKeyByQuotaPeriod() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const quotaCounterKey = "ba";
  const quotaPeriodKey = "0_P3Y6M4DT12H30M5S";
  const parameters = {
    callsCount: 0,
    kbTransferred: 0,
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.quotaByPeriodKeys.update(
    resourceGroupName,
    serviceName,
    quotaCounterKey,
    quotaPeriodKey,
    parameters,
  );
  console.log(result);
}
