const createComputeManagementClient = require("@azure-rest/arm-compute").default,
  { getLongRunningPoller } = require("@azure-rest/arm-compute");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv").config();

/**
 * This sample demonstrates how to Create or update a VM scale set.
 *
 * @summary Create or update a VM scale set.
 * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2022-08-01/examples/virtualMachineScaleSetExamples/VirtualMachineScaleSet_Create_WithSshAuthentication.json
 */
async function createAScaleSetWithSshAuthentication() {
  const credential = new DefaultAzureCredential();
  const client = createComputeManagementClient(credential);
  const subscriptionId = "";
  const resourceGroupName = "myResourceGroup";
  const vmScaleSetName = "{vmss-name}";
  const options = {
    body: {
      location: "westus",
      properties: {
        overprovision: true,
        upgradePolicy: { mode: "Manual" },
        virtualMachineProfile: {
          networkProfile: {
            networkInterfaceConfigurations: [
              {
                name: "{vmss-name}",
                properties: {
                  enableIPForwarding: true,
                  ipConfigurations: [
                    {
                      name: "{vmss-name}",
                      properties: {
                        subnet: {
                          id: "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Network/virtualNetworks/{existing-virtual-network-name}/subnets/{existing-subnet-name}",
                        },
                      },
                    },
                  ],
                  primary: true,
                },
              },
            ],
          },
          osProfile: {
            adminUsername: "{your-username}",
            computerNamePrefix: "{vmss-name}",
            linuxConfiguration: {
              disablePasswordAuthentication: true,
              ssh: {
                publicKeys: [
                  {
                    path: "/home/{your-username}/.ssh/authorized_keys",
                    keyData:
                      "ssh-rsa AAAAB3NzaC1yc2EAAAADAQABAAABAQCeClRAk2ipUs/l5voIsDC5q9RI+YSRd1Bvd/O+axgY4WiBzG+4FwJWZm/mLLe5DoOdHQwmU2FrKXZSW4w2sYE70KeWnrFViCOX5MTVvJgPE8ClugNl8RWth/tU849DvM9sT7vFgfVSHcAS2yDRyDlueii+8nF2ym8XWAPltFVCyLHRsyBp5YPqK8JFYIa1eybKsY3hEAxRCA+/7bq8et+Gj3coOsuRmrehav7rE6N12Pb80I6ofa6SM5XNYq4Xk0iYNx7R3kdz0Jj9XgZYWjAHjJmT0gTRoOnt6upOuxK7xI/ykWrllgpXrCPu3Ymz+c+ujaqcxDopnAl2lmf69/J1",
                  },
                ],
              },
            },
          },
          storageProfile: {
            imageReference: {
              offer: "WindowsServer",
              publisher: "MicrosoftWindowsServer",
              sku: "2016-Datacenter",
              version: "latest",
            },
            osDisk: {
              caching: "ReadWrite",
              createOption: "FromImage",
              managedDisk: { storageAccountType: "Standard_LRS" },
            },
          },
        },
      },
      sku: { name: "Standard_D1_v2", capacity: 3, tier: "Standard" },
    },
    queryParameters: { "api-version": "2022-08-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/resourceGroups/{resourceGroupName}/providers/Microsoft.Compute/virtualMachineScaleSets/{vmScaleSetName}",
      subscriptionId,
      resourceGroupName,
      vmScaleSetName
    )
    .put(options);
  const poller = getLongRunningPoller(client, initialResponse);
  const result = await poller.pollUntilDone();
  console.log(result);
}

createAScaleSetWithSshAuthentication().catch(console.error);
