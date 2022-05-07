Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-resources_5.0.1/sdk/resources/arm-resources/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to When you delete a resource group, all of its resources are also deleted. Deleting a resource group deletes all of its template deployments and currently stored operations.
 *
 * @summary When you delete a resource group, all of its resources are also deleted. Deleting a resource group deletes all of its template deployments and currently stored operations.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/ForceDeleteVMsInResourceGroup.json
 */
async function forceDeleteAllTheVirtualMachinesInAResourceGroup() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "my-resource-group";
  const forceDeletionTypes = "Microsoft.Compute/virtualMachines";
  const options = { forceDeletionTypes };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.resourceGroups.beginDeleteAndWait(resourceGroupName, options);
  console.log(result);
}

forceDeleteAllTheVirtualMachinesInAResourceGroup().catch(console.error);
```
