Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-purview_1.0.1/sdk/purview/arm-purview/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { PurviewManagementClient } = require("@azure/arm-purview");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a private endpoint connection
 *
 * @summary Create or update a private endpoint connection
 * x-ms-original-file: specification/purview/resource-manager/Microsoft.Purview/stable/2021-07-01/examples/PrivateEndpointConnections_CreateOrUpdate.json
 */
async function privateEndpointConnectionsCreateOrUpdate() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "SampleResourceGroup";
  const accountName = "account1";
  const privateEndpointConnectionName = "privateEndpointConnection1";
  const request = {
    privateLinkServiceConnectionState: {
      description: "Approved by johndoe@company.com",
      status: "Approved",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new PurviewManagementClient(credential, subscriptionId);
  const result = await client.privateEndpointConnections.beginCreateOrUpdateAndWait(
    resourceGroupName,
    accountName,
    privateEndpointConnectionName,
    request
  );
  console.log(result);
}

privateEndpointConnectionsCreateOrUpdate().catch(console.error);
```
