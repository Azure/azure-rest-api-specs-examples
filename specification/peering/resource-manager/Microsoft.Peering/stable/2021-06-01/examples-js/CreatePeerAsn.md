Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new peer ASN or updates an existing peer ASN with the specified name under the given subscription.
 *
 * @summary Creates a new peer ASN or updates an existing peer ASN with the specified name under the given subscription.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/CreatePeerAsn.json
 */
async function createAPeerAsn() {
  const subscriptionId = "subId";
  const peerAsnName = "peerAsnName";
  const peerAsn = {
    peerAsn: 65000,
    peerContactDetail: [
      { email: "noc@contoso.com", phone: "+1 (234) 567-8999", role: "Noc" },
      { email: "abc@contoso.com", phone: "+1 (234) 567-8900", role: "Policy" },
      {
        email: "xyz@contoso.com",
        phone: "+1 (234) 567-8900",
        role: "Technical",
      },
    ],
    peerName: "Contoso",
  };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peerAsns.createOrUpdate(peerAsnName, peerAsn);
  console.log(result);
}

createAPeerAsn().catch(console.error);
```
