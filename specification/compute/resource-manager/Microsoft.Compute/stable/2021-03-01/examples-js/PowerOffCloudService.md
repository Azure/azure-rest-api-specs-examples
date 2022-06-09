```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Power off the cloud service. Note that resources are still attached and you are getting charged for the resources.
 *
 * @summary Power off the cloud service. Note that resources are still attached and you are getting charged for the resources.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-03-01/examples/PowerOffCloudService.json
 */
async function stopOrPowerOffCloudService() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ConstosoRG";
  const cloudServiceName = "{cs-name}";
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.cloudServices.beginPowerOffAndWait(
    resourceGroupName,
    cloudServiceName
  );
  console.log(result);
}

stopOrPowerOffCloudService().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_19.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.
