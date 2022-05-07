Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-dns_5.0.1/sdk/dns/arm-dns/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all record sets in a DNS zone.
 *
 * @summary Lists all record sets in a DNS zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/ListRecordSetsByZone.json
 */
async function listRecordsetsByZone() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.recordSets.listAllByDnsZone(resourceGroupName, zoneName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listRecordsetsByZone().catch(console.error);
```
