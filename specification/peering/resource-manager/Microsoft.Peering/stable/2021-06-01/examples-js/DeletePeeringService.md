Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function deleteAPeeringService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peeringServices.delete(resourceGroupName, peeringServiceName);
  console.log(result);
}

deleteAPeeringService().catch(console.error);
```
