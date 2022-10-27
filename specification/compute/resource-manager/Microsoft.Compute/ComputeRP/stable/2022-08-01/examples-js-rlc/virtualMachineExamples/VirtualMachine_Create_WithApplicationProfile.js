const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 *
 * @summary The operation to create or update a virtual machine. Please note some properties can be set only during virtual machine creation.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineExamples/VirtualMachine_Create_WithApplicationProfile.json
 */
async function createAVMWithApplicationProfile() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "myResourceGroup";
  const vmName = "myVM";
  const options = {
    body: {
      location: "westus",
      properties: {
        applicationProfile: {
          galleryApplications: [
            {
              configurationReference:
                "https://mystorageaccount.blob.core.windows.net/configurations/settings.config",
              enableAutomaticUpgrade: false,
              order: 1,
              packageReferenceId:
                "/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdb/resourceGroups/myresourceGroupName2/providers/Microsoft.Compute/galleries/myGallery1/applications/MyApplication1/versions/1.0",
              tags: "myTag1",
              treatFailureAsDeploymentFailure: false,
            },
            {
              packageReferenceId:
                "/subscriptions/32c17a9e-aa7b-4ba5-a45b-e324116b6fdg/resourceGroups/myresourceGroupName3/providers/Microsoft.Compute/galleries/myGallery2/applications/MyApplication2/versions/1.1",
            },
          ],
        },
        hardwareProfile: { vmSize: "Standard_D1_v2" },
        networkProfile: {
          networkInterfaces: [
            {
              id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/networkInterfaces/{existing-nic-name}",
              properties: { primary: true },
            },
          ],
        },
        osProfile: {
          adminPassword: "{your-password}",
          adminUsername: "{your-username}",
          computerName: "myVM",
        },
        storageProfile: {
          imageReference: {
            offer: "{image_offer}",
            publisher: "{image_publisher}",
            sku: "{image_sku}",
            version: "latest",
          },
          osDisk: {
            name: "myVMosdisk",
            caching: "ReadWrite",
            createOption: "FromImage",
            managedDisk: { storageAccountType: "Standard_LRS" },
          },
        },
      },
    },
    queryParameters: { "api-version": "2022-08-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachines/{vmName}",
      subscriptionId,
      resourceGroupName,
      vmName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createAVMWithApplicationProfile().catch(console.error);
