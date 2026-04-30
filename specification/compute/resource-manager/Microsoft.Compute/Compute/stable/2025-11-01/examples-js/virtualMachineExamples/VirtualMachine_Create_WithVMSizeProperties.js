const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Create_WithVMSizeProperties.json
 */
async function createAVMWithVMSizeProperties() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus",
    hardwareProfile: {
      vmSize: "Standard_D4_v3",
      vmSizeProperties: { vCPUsAvailable: 1, vCPUsPerCore: 1 },
    },
    storageProfile: {
      imageReference: {
        sku: "2016-Datacenter",
        publisher: "MicrosoftWindowsServer",
        version: "latest",
        offer: "WindowsServer",
      },
      osDisk: {
        caching: "ReadWrite",
        managedDisk: { storageAccountType: "Standard_LRS" },
        name: "myVMosdisk",
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
      computerName: "myVM",
      adminPassword: "{your-password}",
    },
    diagnosticsProfile: {
      bootDiagnostics: {
        storageUri: "http://{existing-storage-account-name}.blob.core.windows.net",
        enabled: true,
      },
    },
    userData: "U29tZSBDdXN0b20gRGF0YQ==",
  });
  console.log(result);
}
