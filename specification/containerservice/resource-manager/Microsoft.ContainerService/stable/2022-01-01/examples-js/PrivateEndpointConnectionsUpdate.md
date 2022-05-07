Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-containerservice_15.2.0/sdk/containerservice/arm-containerservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerServiceClient } = require("@azure/arm-containerservice");
const { DefaultAzureCredential } = require("@azure/identity");

async function updatePrivateEndpointConnection() {
  const subscriptionId = "subid1";
  const resourceGroupName = "rg1";
  const resourceName = "clustername1";
  const privateEndpointConnectionName = "privateendpointconnection1";
  const parameters = {
    privateLinkServiceConnectionState: { status: "Approved" },
  };
  const credential = new DefaultAzureCredential();
  const client = new ContainerServiceClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.update(
    resourceGroupName,
    resourceName,
    privateEndpointConnectionName,
    parameters
  );
  console.log(result);
}

updatePrivateEndpointConnection().catch(console.error);
```
