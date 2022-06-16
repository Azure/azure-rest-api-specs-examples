```javascript
const { NetAppManagementClient } = require("@azure/arm-netapp");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create the specified quota rule within the given volume
 *
 * @summary Create the specified quota rule within the given volume
 * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/VolumeQuotaRules_Create.json
 */
async function volumeQuotaRulesCreate() {
  const subscriptionId = "5275316f-a498-48d6-b324-2cbfdc4311b9";
  const resourceGroupName = "myRG";
  const accountName = "account-9957";
  const poolName = "pool-5210";
  const volumeName = "volume-6387";
  const volumeQuotaRuleName = "rule-0004";
  const body = {
    location: "westus",
    quotaSizeInKiBs: 100005,
    quotaTarget: "1821",
    quotaType: "IndividualUserQuota",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetAppManagementClient(credential, subscriptionId);
  const result = await client.volumeQuotaRules.beginCreateAndWait(
    resourceGroupName,
    accountName,
    poolName,
    volumeName,
    volumeQuotaRuleName,
    body
  );
  console.log(result);
}

volumeQuotaRulesCreate().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-netapp_16.0.0/sdk/netapp/arm-netapp/README.md) on how to add the SDK to your project and authenticate.
