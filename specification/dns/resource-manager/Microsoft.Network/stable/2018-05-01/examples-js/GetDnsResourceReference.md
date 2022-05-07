Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-dns_5.0.1/sdk/dns/arm-dns/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the DNS records specified by the referencing targetResourceIds.
 *
 * @summary Returns the DNS records specified by the referencing targetResourceIds.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetDnsResourceReference.json
 */
async function listZonesByResourceGroup() {
  const subscriptionId = "subid";
  const parameters = {
    targetResources: [
      {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/trafficManagerProfiles/testpp2",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.dnsResourceReferenceOperations.getByTargetResources(parameters);
  console.log(result);
}

listZonesByResourceGroup().catch(console.error);
```
