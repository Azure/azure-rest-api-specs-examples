const { DataLakeAnalyticsAccountManagementClient } = require("@azure/arm-datalake-analytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the Data Lake Analytics account specified to remove the specified Data Lake Store account.
 *
 * @summary Updates the Data Lake Analytics account specified to remove the specified Data Lake Store account.
 * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/DataLakeStoreAccounts_Delete.json
 */
async function removesTheSpecifiedDataLakeStoreAccount() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "contosorg";
  const accountName = "contosoadla";
  const dataLakeStoreAccountName = "test_adls_account";
  const credential = new DefaultAzureCredential();
  const client = new DataLakeAnalyticsAccountManagementClient(credential, subscriptionId);
  const result = await client.dataLakeStoreAccounts.delete(
    resourceGroupName,
    accountName,
    dataLakeStoreAccountName
  );
  console.log(result);
}

removesTheSpecifiedDataLakeStoreAccount().catch(console.error);
