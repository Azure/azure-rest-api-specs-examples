const { DataLakeAnalyticsAccountManagementClient } = require("@azure/arm-datalake-analytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets details of the specified Data Lake Analytics account.
 *
 * @summary Gets details of the specified Data Lake Analytics account.
 * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/Accounts_Get.json
 */
async function getsDetailsOfTheSpecifiedDataLakeAnalyticsAccount() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "contosorg";
  const accountName = "contosoadla";
  const credential = new DefaultAzureCredential();
  const client = new DataLakeAnalyticsAccountManagementClient(credential, subscriptionId);
  const result = await client.accounts.get(resourceGroupName, accountName);
  console.log(result);
}

getsDetailsOfTheSpecifiedDataLakeAnalyticsAccount().catch(console.error);
