const { DataLakeAnalyticsAccountManagementClient } = require("@azure/arm-datalake-analytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified compute policy from the specified Data Lake Analytics account
 *
 * @summary Deletes the specified compute policy from the specified Data Lake Analytics account
 * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/ComputePolicies_Delete.json
 */
async function deletesTheSpecifiedComputePolicyFromTheAdlaAccount() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "contosorg";
  const accountName = "contosoadla";
  const computePolicyName = "test_policy";
  const credential = new DefaultAzureCredential();
  const client = new DataLakeAnalyticsAccountManagementClient(credential, subscriptionId);
  const result = await client.computePolicies.delete(
    resourceGroupName,
    accountName,
    computePolicyName
  );
  console.log(result);
}

deletesTheSpecifiedComputePolicyFromTheAdlaAccount().catch(console.error);
