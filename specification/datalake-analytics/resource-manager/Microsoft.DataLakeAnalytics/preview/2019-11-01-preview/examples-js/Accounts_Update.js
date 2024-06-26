const { DataLakeAnalyticsAccountManagementClient } = require("@azure/arm-datalake-analytics");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates the Data Lake Analytics account object specified by the accountName with the contents of the account object.
 *
 * @summary Updates the Data Lake Analytics account object specified by the accountName with the contents of the account object.
 * x-ms-original-file: specification/datalake-analytics/resource-manager/Microsoft.DataLakeAnalytics/preview/2019-11-01-preview/examples/Accounts_Update.json
 */
async function updatesTheDataLakeAnalyticsAccountObjectSpecifiedByTheAccountNameWithTheContentsOfTheAccountObject() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "contosorg";
  const accountName = "contosoadla";
  const parameters = {
    computePolicies: [
      {
        name: "test_policy",
        maxDegreeOfParallelismPerJob: 1,
        minPriorityPerJob: 1,
        objectId: "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345",
        objectType: "User",
      },
    ],
    firewallAllowAzureIps: "Enabled",
    firewallRules: [{ name: "test_rule", endIpAddress: "2.2.2.2", startIpAddress: "1.1.1.1" }],
    firewallState: "Enabled",
    maxDegreeOfParallelism: 1,
    maxDegreeOfParallelismPerJob: 1,
    maxJobCount: 1,
    minPriorityPerJob: 1,
    newTier: "Consumption",
    queryStoreRetention: 1,
    tags: { testKey: "test_value" },
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new DataLakeAnalyticsAccountManagementClient(credential, subscriptionId);
  const result = await client.accounts.beginUpdateAndWait(resourceGroupName, accountName, options);
  console.log(result);
}

updatesTheDataLakeAnalyticsAccountObjectSpecifiedByTheAccountNameWithTheContentsOfTheAccountObject().catch(
  console.error
);
