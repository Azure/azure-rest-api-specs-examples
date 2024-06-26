const { DataLakeAnalyticsAccountManagementClient } = require("@azure/arm-datalake-analytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Data Lake Store account details in the specified Data Lake Analytics account.
 *
 * @summary Gets the specified Data Lake Store account details in the specified Data Lake Analytics account.
 * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/DataLakeStoreAccounts_Get.json
 */
async function getsTheSpecifiedDataLakeStoreAccountDetails() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1669ab345";
  const resourceGroupName = "contosorg";
  const accountName = "contosoadla";
  const dataLakeStoreAccountName = "test_adls_account";
  const credential = new DefaultAzureCredential();
  const client = new DataLakeAnalyticsAccountManagementClient(credential, subscriptionId);
  const result = await client.dataLakeStoreAccounts.get(
    resourceGroupName,
    accountName,
    dataLakeStoreAccountName
  );
  console.log(result);
}

getsTheSpecifiedDataLakeStoreAccountDetails().catch(console.error);
