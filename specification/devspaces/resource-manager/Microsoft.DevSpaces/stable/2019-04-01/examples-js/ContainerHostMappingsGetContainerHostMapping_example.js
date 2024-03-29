const { DevSpacesManagementClient } = require("@azure/arm-devspaces");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns container host mapping object for a container host resource ID if an associated controller exists.
 *
 * @summary Returns container host mapping object for a container host resource ID if an associated controller exists.
 * x-ms-original-file: specification/devspaces/resource-manager/Microsoft.DevSpaces/stable/2019-04-01/examples/ContainerHostMappingsGetContainerHostMapping_example.json
 */
async function containerHostMappingsGetContainerHostMapping() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const location = "eastus";
  const containerHostMapping = {
    containerHostResourceId:
      "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/myResourceGroup/providers/Microsoft.ContainerService/managedClusters/myCluster",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevSpacesManagementClient(credential, subscriptionId);
  const result = await client.containerHostMappings.getContainerHostMapping(
    resourceGroupName,
    location,
    containerHostMapping
  );
  console.log(result);
}

containerHostMappingsGetContainerHostMapping().catch(console.error);
