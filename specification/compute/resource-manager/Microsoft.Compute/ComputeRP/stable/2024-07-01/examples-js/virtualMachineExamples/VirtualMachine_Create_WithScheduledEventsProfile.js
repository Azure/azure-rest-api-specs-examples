const { ComputeManagementClient } = require("@azure/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/virtualMachineExamples/VirtualMachine_Create_WithScheduledEventsProfile.json
 */
async function createAVMWithScheduledEventsProfile() {
  const subscriptionId = process.env["COMPUTE_SUBSCRIPTION_ID"] || "{subscription-id}";
  const resourceGroupName = process.env["COMPUTE_RESOURCE_GROUP"] || "myResourceGroup";
  const vmName = "myVM";
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
      computerName: "myVM",
    },
    scheduledEventsPolicy: {
      scheduledEventsAdditionalPublishingTargets: {
        eventGridAndResourceGraph: { enable: true },
      },
      userInitiatedReboot: { automaticallyApprove: true },
      userInitiatedRedeploy: { automaticallyApprove: true },
    },
    scheduledEventsProfile: {
      osImageNotificationProfile: { enable: true, notBeforeTimeout: "PT15M" },
      terminateNotificationProfile: { enable: true, notBeforeTimeout: "PT10M" },
    },
    storageProfile: {
      imageReference: {
        offer: "WindowsServer",
        publisher: "MicrosoftWindowsServer",
        sku: "2016-Datacenter",
        version: "latest",
      },
      osDisk: {
        name: "myVMosdisk",
        caching: "ReadWrite",
        createOption: "FromImage",
        managedDisk: { storageAccountType: "Standard_LRS" },
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
