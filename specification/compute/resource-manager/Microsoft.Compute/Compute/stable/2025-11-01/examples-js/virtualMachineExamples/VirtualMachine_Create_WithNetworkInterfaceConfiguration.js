const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary the operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: 2025-11-01/virtualMachineExamples/VirtualMachine_Create_WithNetworkInterfaceConfiguration.json
 */
async function createAVMWithNetworkInterfaceConfiguration() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "{subscription-id}";
  const client = new ComputeManagementClient(credential, subscriptionId);
  const result = await client.virtualMachines.createOrUpdate("myResourceGroup", "myVM", {
    location: "westus",
    hardwareProfile: { vmSize: "Standard_D1_v2" },
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
      networkApiVersion: "2020-11-01",
      networkInterfaceConfigurations: [
        {
          name: "{nic-config-name}",
          tags: { nicTag: "tag" },
          primary: true,
          deleteOption: "Delete",
          ipConfigurations: [
            {
              name: "{ip-config-name}",
              primary: true,
              publicIPAddressConfiguration: {
                name: "{publicIP-config-name}",
                tags: { pipTag: "tag" },
                sku: { name: "Basic", tier: "Global" },
                deleteOption: "Detach",
                publicIPAllocationMethod: "Static",
              },
            },
          ],
        },
      ],
    },
    osProfile: {
      adminUsername: "{your-username}",
      computerName: "myVM",
      adminPassword: "{your-password}",
    },
  });
  console.log(result);
}
