const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Create_WithHibernationEnabled.json
 */
async function createAVMWithHibernationEnabled() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "{vm-name}", {
    location: "eastus2euap",
    hardwareProfile: { vmSize: "Standard_D2s_v3" },
    additionalCapabilities: { hibernationEnabled: true },
    storageProfile: {
      imageReference: {
        publisher: "MicrosoftWindowsServer",
        offer: "WindowsServer",
        sku: "2019-Datacenter",
        version: "latest",
      },
      osDisk: {
        caching: "ReadWrite",
        managedDisk: { storageAccountType: "Standard_LRS" },
        name: "vmOSdisk",
        createOption: "FromImage",
      },
    },
    networkProfile: {
      networkInterfaces: [
        {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
          primary: true,
        },
      ],
    },
    osProfile: {
      adminUsername: "{your-username}",
      computerName: "{vm-name}",
      adminPassword: "{your-password}",
    },
    diagnosticsProfile: {
      bootDiagnostics: {
        storageUri: "http://{existing-storage-account-name}.blob.core.windows.net",
        enabled: true,
      },
    },
  });
  console.log(result);
}
