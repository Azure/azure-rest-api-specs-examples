const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2023-09-01/examples/virtualMachineExamples/VirtualMachine_Create_WithEncryptionIdentity.json
 */
async function createAVMWithEncryptionIdentity() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmName = "myVM";
  const parameters = {
    hardwareProfile: { vmSize: "Standard_D2s_v3" },
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/MicrosoftManagedIdentity/userAssignedIdentities/myIdentity":
          {},
      },
    },
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
      encryptionIdentity: {
        userAssignedIdentityResourceId:
          "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.ManagedIdentity/userAssignedIdentities/myIdentity",
      },
    },
    storageProfile: {
      imageReference: {
        offer: "WindowsServer",
        publisher: "MicrosoftWindowsServer",
        sku: "2019-Datacenter",
        version: "latest",
      },
      osDisk: {
        name: "myVMosdisk",
        caching: "ReadOnly",
        createOption: "FromImage",
        managedDisk: { storageAccountType: "StandardSSD_LRS" },
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
