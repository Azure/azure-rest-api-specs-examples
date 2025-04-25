const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a Virtual Instance for SAP solutions (VIS) resource
 *
 * @summary creates a Virtual Instance for SAP solutions (VIS) resource
 * x-ms-original-file: 2024-09-01/SapVirtualInstances_CreateInstallSingleServer.json
 */
async function installSAPSoftwareOnSingleServerSystem() {
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
          appResourceGroup: "test-rg",
          deploymentType: "SingleServer",
          subnetId:
            "/subscriptions/8e17e36c-42e9-4cd5-a078-7b44883414e0/resourceGroups/test-rg/providers/Microsoft.Network/virtualNetworks/test-vnet/subnets/testsubnet",
          virtualMachineConfiguration: {
            imageReference: {
              offer: "SLES-SAP",
              publisher: "SUSE",
              sku: "12-sp4-gen2",
              version: "2022.02.01",
            },
            osProfile: {
              adminUsername: "azureappadmin",
              osConfiguration: {
                disablePasswordAuthentication: true,
                osType: "Linux",
                sshKeyPair: {
                  privateKey: "{{privateKey}}",
                  publicKey: "{{sshkey}}",
                },
              },
            },
            vmSize: "Standard_E32ds_v4",
          },
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
      environment: "NonProd",
      sapProduct: "S4HANA",
    },
    tags: {},
  });
  console.log(result);
}
