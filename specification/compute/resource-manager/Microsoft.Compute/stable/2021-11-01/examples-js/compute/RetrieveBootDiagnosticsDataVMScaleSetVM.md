Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function retrieveBootDiagnosticsDataOfAVirtualMachine() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "ResourceGroup";
  const vmScaleSetName = "myvmScaleSet";
  const instanceId = "0";
  const sasUriExpirationTimeInMinutes = 60;
  const options = {
    sasUriExpirationTimeInMinutes: sasUriExpirationTimeInMinutes,
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMs.retrieveBootDiagnosticsData(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
    options
  );
  console.log(result);
}

retrieveBootDiagnosticsDataOfAVirtualMachine().catch(console.error);
```
