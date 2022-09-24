const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get details of the specified quota rule
 *
 * @summary Get details of the specified quota rule
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-05-01/examples/VolumeQuotaRules_Get.json
 */
async function volumeQuotaRulesGet() {
  const subscriptionId = "5275316f-a498-48d6-b324-2cbfdc4311b9";
  const resourceGroupName = "myRG";
  const accountName = "account-9957";
  const poolName = "pool-5210";
  const volumeName = "volume-6387";
  const volumeQuotaRuleName = "rule-0004";
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.volumeQuotaRules.get(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    volumeQuotaRuleName
  );
  console.log(result);
}

volumeQuotaRulesGet().catch(console.error);
