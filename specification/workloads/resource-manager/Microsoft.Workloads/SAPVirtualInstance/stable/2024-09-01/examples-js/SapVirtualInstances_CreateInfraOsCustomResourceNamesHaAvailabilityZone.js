const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a Virtual Instance for SAP solutions (VIS) resource
 *
 * @summary creates a Virtual Instance for SAP solutions (VIS) resource
 * x-ms-original-file: 2024-09-01/SapVirtualInstances_CreateInfraOsCustomResourceNamesHaAvailabilityZone.json
 */
async function createInfrastructureWithOSConfigurationWithCustomResourceNamesForHASystemWithAvailabilityZone() {
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
            instanceCount: 2,
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
              loadBalancer: {
                backendPoolNames: ["ascsBackendPool"],
                frontendIpConfigurationNames: ["ascsip0", "ersip0"],
                healthProbeNames: ["ascsHealthProbe", "ersHealthProbe"],
                loadBalancerName: "ascslb",
              },
              virtualMachines: [
                {
                  hostName: "ascshostName",
                  networkInterfaces: [{ networkInterfaceName: "ascsnic" }],
                  osDiskName: "ascsosdisk",
                  vmName: "ascsvm",
                },
                {
                  hostName: "ershostName",
                  networkInterfaces: [{ networkInterfaceName: "ersnic" }],
                  osDiskName: "ersosdisk",
                  vmName: "ersvm",
                },
              ],
            },
            databaseServer: {
              loadBalancer: {
                backendPoolNames: ["dbBackendPool"],
                frontendIpConfigurationNames: ["dbip"],
                healthProbeNames: ["dbHealthProbe"],
                loadBalancerName: "dblb",
              },
              virtualMachines: [
                {
                  dataDiskNames: {
                    hanaData: ["hanadatapr0", "hanadatapr1"],
                    hanaLog: ["hanalogpr0", "hanalogpr1", "hanalogpr2"],
                    hanaShared: ["hanasharedpr0", "hanasharedpr1"],
                    usrSap: ["usrsappr0"],
                  },
                  hostName: "dbprhostName",
                  networkInterfaces: [{ networkInterfaceName: "dbprnic" }],
                  osDiskName: "dbprosdisk",
                  vmName: "dbvmpr",
                },
                {
                  dataDiskNames: {
                    hanaData: ["hanadatasr0", "hanadatasr1"],
                    hanaLog: ["hanalogsr0", "hanalogsr1", "hanalogsr2"],
                    hanaShared: ["hanasharedsr0", "hanasharedsr1"],
                    usrSap: ["usrsapsr0"],
                  },
                  hostName: "dbsrhostName",
                  networkInterfaces: [{ networkInterfaceName: "dbsrnic" }],
                  osDiskName: "dbsrosdisk",
                  vmName: "dbvmsr",
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
            instanceCount: 2,
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
          highAvailabilityConfig: { highAvailabilityType: "AvailabilityZone" },
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
