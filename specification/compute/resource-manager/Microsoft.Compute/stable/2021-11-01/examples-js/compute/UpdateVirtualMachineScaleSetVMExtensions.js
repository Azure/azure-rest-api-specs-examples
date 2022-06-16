const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

async function updateVirtualMachineScaleSetVMExtension() {
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
  const result = await client.virtualMachineScaleSetVMExtensions.beginUpdateAndWait(
    resourceGroupName,
    vmScaleSetName,
    instanceId,
    vmExtensionName,
    extensionParameters
  );
  console.log(result);
}

updateVirtualMachineScaleSetVMExtension().catch(console.error);
