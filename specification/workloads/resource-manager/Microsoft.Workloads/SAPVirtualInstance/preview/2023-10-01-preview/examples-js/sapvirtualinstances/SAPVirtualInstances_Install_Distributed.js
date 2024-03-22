const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Virtual Instance for SAP solutions (VIS) resource
 *
 * @summary Creates a Virtual Instance for SAP solutions (VIS) resource
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Install_Distributed.json
 */
async function installSapSoftwareOnDistributedSystem() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const body = {
    location: "eastus2",
    properties: {
      configuration: {
        appLocation: "eastus",
        configurationType: "DeploymentWithOSConfig",
        infrastructureConfiguration: {
          appResourceGroup: "{{resourcegrp}}",
          applicationServer: {
            instanceCount: 2,
            subnetId:
              "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app",
            virtualMachineConfiguration: {
              imageReference: {
                offer: "RHEL-SAP-HA",
                publisher: "RedHat",
                sku: "8.2",
                version: "8.2.2021091201",
              },
              osProfile: {
                adminUsername: "azureuser",
                osConfiguration: {
                  disablePasswordAuthentication: true,
                  osType: "Linux",
                  sshKeyPair: {
                    privateKey: "{{privateKey}}",
                    publicKey: "{{sshkey}}",
                  },
                },
              },
              vmSize: "Standard_E4ds_v4",
            },
          },
          centralServer: {
            instanceCount: 1,
            subnetId:
              "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app",
            virtualMachineConfiguration: {
              imageReference: {
                offer: "RHEL-SAP-HA",
                publisher: "RedHat",
                sku: "8.2",
                version: "8.2.2021091201",
              },
              osProfile: {
                adminUsername: "azureuser",
                osConfiguration: {
                  disablePasswordAuthentication: true,
                  osType: "Linux",
                  sshKeyPair: {
                    privateKey: "{{privateKey}}",
                    publicKey: "{{sshkey}}",
                  },
                },
              },
              vmSize: "Standard_E4ds_v4",
            },
          },
          databaseServer: {
            instanceCount: 1,
            subnetId:
              "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app",
            virtualMachineConfiguration: {
              imageReference: {
                offer: "RHEL-SAP-HA",
                publisher: "RedHat",
                sku: "8.2",
                version: "8.2.2021091201",
              },
              osProfile: {
                adminUsername: "azureuser",
                osConfiguration: {
                  disablePasswordAuthentication: true,
                  osType: "Linux",
                  sshKeyPair: {
                    privateKey: "{{privateKey}}",
                    publicKey: "{{sshkey}}",
                  },
                },
              },
              vmSize: "Standard_M32ts",
            },
          },
          deploymentType: "ThreeTier",
          networkConfiguration: { isSecondaryIpEnabled: true },
        },
        osSapConfiguration: { sapFqdn: "sap.bpaas.com" },
        softwareConfiguration: {
          bomUrl:
            "https://teststorageaccount.blob.core.windows.net/sapbits/sapfiles/boms/S41909SPS03_v0011ms/S41909SPS03_v0011ms.yaml",
          sapBitsStorageAccountId:
            "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Storage/storageAccounts/teststorageaccount",
          softwareInstallationType: "SAPInstallWithoutOSConfig",
          softwareVersion: "SAP S/4HANA 1909 SPS 03",
        },
      },
      environment: "Prod",
      sapProduct: "S4HANA",
    },
    tags: { createdBy: "azureuser" },
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPVirtualInstances.beginCreateAndWait(
    resourceGroupName,
    sapVirtualInstanceName,
    options,
  );
  console.log(result);
}
