const { DataLakeAnalyticsAccountManagementClient } = require("@azure/arm-datalake-analytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the specified Data Lake Analytics account to include the additional Data Lake Store account.
 *
 * @summary Updates the specified Data Lake Analytics account to include the additional Data Lake Store account.
 * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/DataLakeStoreAccounts_Add.json
 */
async function addsADataLakeStoreAccount() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "contosorg";
  const accountName = "contosoadla";
  const dataLakeStoreAccountName = "test_adls_account";
  const parameters = { suffix: "test_suffix" };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new DataLakeAnalyticsAccountManagementClient(credential, subscriptionId);
  const result = await client.dataLakeStoreAccounts.add(
    resourceGroupName,
    accountName,
    dataLakeStoreAccountName,
    options
  );
  console.log(result);
}

addsADataLakeStoreAccount().catch(console.error);
