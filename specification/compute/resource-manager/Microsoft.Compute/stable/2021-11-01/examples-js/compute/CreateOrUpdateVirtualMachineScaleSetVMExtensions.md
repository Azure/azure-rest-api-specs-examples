```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update the VMSS VM extension.
 *
 * @summary The operation to create or update the VMSS VM extension.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-11-01/examples/compute/CreateOrUpdateVirtualMachineScaleSetVMExtensions.json
 */
async function createVirtualMachineScaleSetVMExtension() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "myvmScaleSet";
  const instanceId = "0";
  const vmExtensionName = "myVMExtension";
  const extensionParameters = {
    typePropertiesType: "extType",
    autoUpgradeMinorVersion: true,
    publisher: "extPublisher",
    settings: { UserName: "xyz@microsoft.com" },
    typeHandlerVersion: "1.2",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineScaleSetVMExtensions.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
    vmExtensionName,
    extensionParameters
  );
  console.log(result);
}

createVirtualMachineScaleSetVMExtension().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_17.3.1/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.
