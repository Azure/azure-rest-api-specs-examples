const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Virtual Instance for SAP solutions (VIS) resource
 *
 * @summary Creates a Virtual Instance for SAP solutions (VIS) resource
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Create_CustomFullResourceNames_Distributed.json
 */
async function createInfrastructureWithOSConfigurationWithCustomResourceNamesForDistributedSystem() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const body = {
    location: "eastus",
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
      managedResourceGroupConfiguration: {
        name: "mrg-X00-8e17e36c-42e9-4cd5-a078-7b44883414e0",
      },
      sapProduct: "S4HANA",
    },
    tags: {},
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
