```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function listAllTheCdnPeeringPrefixesAdvertisedAtAParticularPeeringLocation() {
  const subscriptionId = "subId";
  const peeringLocation = "peeringLocation0";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.cdnPeeringPrefixes.list(peeringLocation)) {
    resArray.push(item);
  }
  console.log(resArray);
}

listAllTheCdnPeeringPrefixesAdvertisedAtAParticularPeeringLocation().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.
