Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-peering_2.0.1/sdk/peering/arm-peering/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

async function createAPeeringService() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringServiceName = "peeringServiceName";
  const peeringService = {
    location: "eastus",
    peeringServiceLocation: "state1",
    peeringServiceProvider: "serviceProvider1",
    providerBackupPeeringLocation: "peeringLocation2",
    providerPrimaryPeeringLocation: "peeringLocation1",
  };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const result = await client.peeringServices.createOrUpdate(
    resourceGroupName,
    peeringServiceName,
    peeringService
  );
  console.log(result);
}

createAPeeringService().catch(console.error);
```
