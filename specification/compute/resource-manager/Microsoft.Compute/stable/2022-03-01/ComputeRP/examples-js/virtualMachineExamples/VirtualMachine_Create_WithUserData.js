const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2022-03-01/ComputeRP/examples/virtualMachineExamples/VirtualMachine_Create_WithUserData.json
 */
async function createAVMWithUserData() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "myResourceGroup";
  const vmName = "{vm-name}";
  const parameters = {
    diagnosticsProfile: {
      bootDiagnostics: {
        enabled: true,
        storageUri: "http://{existing-storage-account-name}.blob.core.windows.net",
      },
    },
    hardwareProfile: { vmSize: "Standard_D1_v2" },
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
      computerName: "{vm-name}",
    },
    storageProfile: {
      imageReference: {
        offer: "WindowsServer",
        publisher: "MicrosoftWindowsServer",
        sku: "2016-Datacenter",
        version: "latest",
      },
      osDisk: {
        name: "vmOSdisk",
        caching: "ReadWrite",
        createOption: "FromImage",
        managedDisk: { storageAccountType: "Standard_LRS" },
      },
    },
    userData: "RXhhbXBsZSBVc2VyRGF0YQ==",
  };
  const credential = new DefaultAzureCredential();
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginCreateOrUpdateAndWait(
    resourceGroupName,
    vmName,
    parameters
  );
  console.log(result);
}

createAVMWithUserData().catch(console.error);
