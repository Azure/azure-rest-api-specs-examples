const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates a Virtual Instance for SAP solutions (VIS) resource
 *
 * @summary creates a Virtual Instance for SAP solutions (VIS) resource
 * x-ms-original-file: 2024-09-01/SapVirtualInstances_CreateInfraHaAvailabilitySet.json
 */
async function createInfrastructureOnlyForHASystemWithAvailabilitySet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sapVirtualInstances.create("test-rg", "X00", {
    location: "westcentralus",
    properties: {
      configuration: {
        appLocation: "eastus",
        configurationType: "Deployment",
        infrastructureConfiguration: {
          appResourceGroup: "X00-RG",
          applicationServer: {
            instanceCount: 5,
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
                  ssh: { publicKeys: [{ keyData: "ssh-rsa public key" }] },
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
                  ssh: { publicKeys: [{ keyData: "ssh-rsa public key" }] },
                },
              },
              vmSize: "Standard_E16ds_v4",
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
                  ssh: { publicKeys: [{ keyData: "ssh-rsa public key" }] },
                },
              },
              vmSize: "Standard_M32ts",
            },
          },
          deploymentType: "ThreeTier",
          highAvailabilityConfig: { highAvailabilityType: "AvailabilitySet" },
        },
      },
      environment: "Prod",
      sapProduct: "S4HANA",
    },
    tags: {},
  });
  console.log(result);
}
