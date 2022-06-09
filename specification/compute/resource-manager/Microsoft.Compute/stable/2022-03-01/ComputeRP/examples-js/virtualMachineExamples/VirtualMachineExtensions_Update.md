Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-compute_19.0.0/sdk/compute/arm-compute/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to update the extension.
 *
 * @summary The operation to update the extension.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachineExtensions_Update.json
 */
async function updateVMExtension() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVM";
  const vmExtensionName = "myVMExtension";
  const extensionParameters = {
    type: "extType",
    autoUpgradeMinorVersion: true,
    protectedSettingsFromKeyVault: {
      secretUrl:
        "https://kvName.vault.azure.net/secrets/secretName/79b88b3a6f5440ffb2e73e44a0db712e",
      sourceVault: {
        id: "/subscriptions/a53f7094-a16c-47af-abe4-b05c05d0d79a/resourceGroups/myResourceGroup/providers/Microsoft.KeyVault/vaults/kvName",
      },
    },
    publisher: "extPublisher",
    settings: { UserName: "xyz@microsoft.com" },
    suppressFailures: true,
    typeHandlerVersion: "1.2",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachineExtensions.beginUpdateAndWait(
    resourceGroupName,
    vmName,
    vmExtensionName,
    extensionParameters
  );
  console.log(result);
}

updateVMExtension().catch(console.error);
```
