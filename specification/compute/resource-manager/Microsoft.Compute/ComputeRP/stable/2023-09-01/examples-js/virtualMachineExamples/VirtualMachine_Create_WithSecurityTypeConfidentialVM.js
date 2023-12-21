const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineExamples/VirtualMachine_Create_WithSecurityTypeConfidentialVM.json
 */
async function createAVMWithSecurityTypeConfidentialVMWithPlatformManagedKeys() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmName = "myVM";
  const parameters = {
    hardwareProfile: { vmSize: "Standard_DC2as_v5" },
    location: "westus",
    networkProfile: {
      networkInterfaces: [
        {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
          primary: true,
        },
      ],
    },
    osProfile: {
      adminPassword: "{your-password}",
      adminUsername: "{your-username}",
      computerName: "myVM",
    },
    securityProfile: {
      securityType: "ConfidentialVM",
      uefiSettings: { secureBootEnabled: true, vTpmEnabled: true },
    },
    storageProfile: {
      imageReference: {
        offer: "2019-datacenter-cvm",
        publisher: "MicrosoftWindowsServer",
        sku: "windows-cvm",
        version: "17763.2183.2109130127",
      },
      osDisk: {
        name: "myVMosdisk",
        caching: "ReadOnly",
        createOption: "FromImage",
        managedDisk: {
          securityProfile: { securityEncryptionType: "DiskWithVMGuestState" },
          storageAccountType: "StandardSSD_LRS",
        },
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmName,
    parameters,
  );
  console.log(result);
}
