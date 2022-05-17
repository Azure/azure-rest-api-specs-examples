Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.5.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function privateEndPointConnectionsListByFactory() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateEndPointConnections.listByFactory(
    resourceGroupName,
    factoryName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

privateEndPointConnectionsListByFactory().catch(console.error);
```
