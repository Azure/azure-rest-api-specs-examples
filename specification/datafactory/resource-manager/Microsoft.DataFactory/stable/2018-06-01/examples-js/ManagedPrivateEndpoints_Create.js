const { DataFactoryManagementClient } = require("@azure/arm-datafactory");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a managed private endpoint.
 *
 * @summary Creates or updates a managed private endpoint.
 * x-ms-original-file: specification/datafactory/resource-manager/Microsoft.DataFactory/stable/2018-06-01/examples/ManagedPrivateEndpoints_Create.json
 */
async function managedVirtualNetworksCreate() {
  const subscriptionId = "12345678-1234-1234-1234-12345678abc";
  const resourceGroupName = "exampleResourceGroup";
  const factoryName = "exampleFactoryName";
  const managedVirtualNetworkName = "exampleManagedVirtualNetworkName";
  const managedPrivateEndpointName = "exampleManagedPrivateEndpointName";
  const managedPrivateEndpoint = {
    properties: {
      fqdns: [],
      groupId: "blob",
      privateLinkResourceId:
        "/subscriptions/12345678-1234-1234-1234-12345678abc/resourceGroups/exampleResourceGroup/providers/Microsoft.Storage/storageAccounts/exampleBlobStorage",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DataFactoryManagementClient(credential, subscriptionId);
  const result = await client.managedPrivateEndpoints.createOrUpdate(
    resourceGroupName,
    factoryName,
    managedVirtualNetworkName,
    managedPrivateEndpointName,
    managedPrivateEndpoint
  );
  console.log(result);
}

managedVirtualNetworksCreate().catch(console.error);
