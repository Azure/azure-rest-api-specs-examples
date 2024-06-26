const { DataLakeAnalyticsAccountManagementClient } = require("@azure/arm-datalake-analytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the specified compute policy.
 *
 * @summary Updates the specified compute policy.
 * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/ComputePolicies_Update.json
 */
async function updatesTheSpecifiedComputePolicy() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "contosorg";
  const accountName = "contosoadla";
  const computePolicyName = "test_policy";
  const parameters = {
    maxDegreeOfParallelismPerJob: 11,
    minPriorityPerJob: 31,
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new DataLakeAnalyticsAccountManagementClient(credential, subscriptionId);
  const result = await client.computePolicies.update(
    resourceGroupName,
    accountName,
    computePolicyName,
    options
  );
  console.log(result);
}

updatesTheSpecifiedComputePolicy().catch(console.error);
