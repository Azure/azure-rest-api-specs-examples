Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-dns_5.0.1/sdk/dns/arm-dns/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the DNS zones in all resource groups in a subscription.
 *
 * @summary Lists the DNS zones in all resource groups in a subscription.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListZonesBySubscription.json
 */
async function listZonesBySubscription() {
  const subscriptionId = "subid";
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.zones.list()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listZonesBySubscription().catch(console.error);
```
