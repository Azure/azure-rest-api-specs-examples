Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datafactory_10.5.0/sdk/datafactory/arm-datafactory/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

async function approvesOrRejectsAPrivateEndpointConnectionForAFactory() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const privateEndpointConnectionName = "connection";
  const privateEndpointWrapper = {
    properties: {
      privateLinkServiceConnectionState: {
        description: "Approved by admin.",
        actionsRequired: "",
        status: "Approved",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnection.createOrUpdate(
    resourceGroupName,
    factoryName,
    privateEndpointConnectionName,
    privateEndpointWrapper
  );
  console.log(result);
}

approvesOrRejectsAPrivateEndpointConnectionForAFactory().catch(console.error);
```
