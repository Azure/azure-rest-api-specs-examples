const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a Virtual Instance for SAP solutions (VIS) resource
 *
 * @summary creates a Virtual Instance for SAP solutions (VIS) resource
 * x-ms-original-file: 2024-09-01/SapVirtualInstances_CreateDetectDS.json
 */
async function detectSAPSoftwareInstallationOnADistributedSystem() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapVirtualInstances.create("test-rg", "X00", {
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
              "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app",
            virtualMachineConfiguration: {
              imageReference: {
                offer: "RHEL-SAP-HA",
                publisher: "RedHat",
                sku: "84sapha-gen2",
                version: "latest",
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
              "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app",
            virtualMachineConfiguration: {
              imageReference: {
                offer: "RHEL-SAP-HA",
                publisher: "RedHat",
                sku: "84sapha-gen2",
                version: "latest",
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
              "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/app",
            virtualMachineConfiguration: {
              imageReference: {
                offer: "RHEL-SAP-HA",
                publisher: "RedHat",
                sku: "84sapha-gen2",
                version: "latest",
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
          centralServerVmId:
            "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Compute/virtualMachines/sapq20scsvm0",
          softwareInstallationType: "External",
        },
      },
      environment: "Prod",
      sapProduct: "S4HANA",
    },
    tags: { "created by": "azureuser" },
  });
  console.log(result);
}
