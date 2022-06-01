Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_18.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Reimage Role Instance asynchronous operation reinstalls the operating system on instances of web roles or worker roles.
 *
 * @summary The Reimage Role Instance asynchronous operation reinstalls the operating system on instances of web roles or worker roles.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/ReimageCloudServiceRoleInstance.json
 */
async function reimageCloudServiceRoleInstance() {
  const subscriptionId = "{subscription-id}";
  const roleInstanceName = "{roleInstance-name}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServiceRoleInstances.beginReimageAndWait(
    roleInstanceName,
    resourceGroupName,
    cloudServiceName
  );
  console.log(result);
}

reimageCloudServiceRoleInstance().catch(console.error);
```
