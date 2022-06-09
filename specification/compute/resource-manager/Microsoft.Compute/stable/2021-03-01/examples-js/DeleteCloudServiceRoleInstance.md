Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_19.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a role instance from a cloud service.
 *
 * @summary Deletes a role instance from a cloud service.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/DeleteCloudServiceRoleInstance.json
 */
async function deleteCloudServiceRoleInstance() {
  const subscriptionId = "{subscription-id}";
  const roleInstanceName = "{roleInstance-name}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServiceRoleInstances.beginDeleteAndWait(
    roleInstanceName,
    resourceGroupName,
    cloudServiceName
  );
  console.log(result);
}

deleteCloudServiceRoleInstance().catch(console.error);
```
