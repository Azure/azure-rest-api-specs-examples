const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Create_WithEncryptionIdentity.json
 */
async function createAVMWithEncryptionIdentity() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus",
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity":
          {},
      },
    },
    hardwareProfile: { vmSize: "Standard_D2s_v3" },
    securityProfile: {
      encryptionIdentity: {
        userAssignedIdentityResourceId:
          "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity",
      },
    },
    storageProfile: {
      imageReference: {
        publisher: "MicrosoftWindowsServer",
        offer: "WindowsServer",
        sku: "2019-Datacenter",
        version: "latest",
      },
      osDisk: {
        caching: "ReadOnly",
        managedDisk: { storageAccountType: "StandardSSD_LRS" },
        createOption: "FromImage",
        name: "myVMosdisk",
      },
    },
    osProfile: {
      adminUsername: "{your-username}",
      computerName: "myVM",
      adminPassword: "{your-password}",
    },
    networkProfile: {
      networkInterfaces: [
        {
          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
          primary: true,
        },
      ],
    },
  });
  console.log(result);
}
