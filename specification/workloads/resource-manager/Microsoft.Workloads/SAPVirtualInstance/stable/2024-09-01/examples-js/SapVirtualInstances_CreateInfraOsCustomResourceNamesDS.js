const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a Virtual Instance for SAP solutions (VIS) resource
 *
 * @summary creates a Virtual Instance for SAP solutions (VIS) resource
 * x-ms-original-file: 2024-09-01/SapVirtualInstances_CreateInfraOsCustomResourceNamesDS.json
 */
async function createInfrastructureWithOSConfigurationWithCustomResourceNamesForDistributedSystem() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapVirtualInstances.create("test-rg", "X00", {
    location: "westcentralus",
    properties: {
      configuration: {
        appLocation: "eastus",
        configurationType: "DeploymentWithOSConfig",
        infrastructureConfiguration: {
          appResourceGroup: "X00-RG",
          applicationServer: {
            instanceCount: 6,
            subnetId:
              "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet",
            virtualMachineConfiguration: {
              imageReference: {
                offer: "RHEL-SAP",
                publisher: "RedHat",
                sku: "84sapha-gen2",
                version: "latest",
              },
              osProfile: {
                adminUsername: "{your-username}",
                osConfiguration: {
                  disablePasswordAuthentication: true,
                  osType: "Linux",
                  sshKeyPair: { privateKey: "xyz", publicKey: "abc" },
                },
              },
              vmSize: "Standard_E32ds_v4",
            },
          },
          centralServer: {
            instanceCount: 1,
            subnetId:
              "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet",
            virtualMachineConfiguration: {
              imageReference: {
                offer: "RHEL-SAP",
                publisher: "RedHat",
                sku: "84sapha-gen2",
                version: "latest",
              },
              osProfile: {
                adminUsername: "{your-username}",
                osConfiguration: {
                  disablePasswordAuthentication: true,
                  osType: "Linux",
                  sshKeyPair: { privateKey: "xyz", publicKey: "abc" },
                },
              },
              vmSize: "Standard_E16ds_v4",
            },
          },
          customResourceNames: {
            applicationServer: {
              availabilitySetName: "appAvSet",
              virtualMachines: [
                {
                  dataDiskNames: { default: ["app0disk0"] },
                  hostName: "apphostName0",
                  networkInterfaces: [{ networkInterfaceName: "appnic0" }],
                  osDiskName: "app0osdisk",
                  vmName: "appvm0",
                },
                {
                  dataDiskNames: { default: ["app1disk0"] },
                  hostName: "apphostName1",
                  networkInterfaces: [{ networkInterfaceName: "appnic1" }],
                  osDiskName: "app1osdisk",
                  vmName: "appvm1",
                },
              ],
            },
            centralServer: {
              virtualMachines: [
                {
                  dataDiskNames: { default: ["ascsdisk0"] },
                  hostName: "ascshostName",
                  networkInterfaces: [{ networkInterfaceName: "ascsnic" }],
                  osDiskName: "ascsosdisk",
                  vmName: "ascsvm",
                },
              ],
            },
            databaseServer: {
              virtualMachines: [
                {
                  dataDiskNames: {
                    hanaData: ["hanadata0", "hanadata1"],
                    hanaLog: ["hanalog0", "hanalog1", "hanalog2"],
                    hanaShared: ["hanashared0", "hanashared1"],
                    usrSap: ["usrsap0"],
                  },
                  hostName: "dbhostName",
                  networkInterfaces: [{ networkInterfaceName: "dbnic" }],
                  osDiskName: "dbosdisk",
                  vmName: "dbvm",
                },
              ],
            },
            namingPatternType: "FullResourceName",
            sharedStorage: {
              sharedStorageAccountName: "storageacc",
              sharedStorageAccountPrivateEndPointName: "peForxNFS",
            },
          },
          databaseServer: {
            databaseType: "HANA",
            instanceCount: 1,
            subnetId:
              "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/test-rg/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/dbsubnet",
            virtualMachineConfiguration: {
              imageReference: {
                offer: "RHEL-SAP",
                publisher: "RedHat",
                sku: "84sapha-gen2",
                version: "latest",
              },
              osProfile: {
                adminUsername: "{your-username}",
                osConfiguration: {
                  disablePasswordAuthentication: true,
                  osType: "Linux",
                  sshKeyPair: { privateKey: "xyz", publicKey: "abc" },
                },
              },
              vmSize: "Standard_M32ts",
            },
          },
          deploymentType: "ThreeTier",
        },
        osSapConfiguration: { sapFqdn: "xyz.test.com" },
      },
      environment: "Prod",
      sapProduct: "S4HANA",
    },
    tags: {},
  });
  console.log(result);
}
