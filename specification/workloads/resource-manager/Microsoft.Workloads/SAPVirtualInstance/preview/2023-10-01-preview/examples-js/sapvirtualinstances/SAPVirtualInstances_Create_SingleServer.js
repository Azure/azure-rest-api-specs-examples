const { WorkloadsClient } = require("@azure/arm-workloadssapvirtualinstance");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a Virtual Instance for SAP solutions (VIS) resource
 *
 * @summary Creates a Virtual Instance for SAP solutions (VIS) resource
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/SAPVirtualInstance/preview/2023-10-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Create_SingleServer.json
 */
async function createInfrastructureOnlyForSingleServerSystem() {
  const subscriptionId =
    process.env["WORKLOADS_SUBSCRIPTION_ID"] || "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const resourceGroupName = process.env["WORKLOADS_RESOURCE_GROUP"] || "test-rg";
  const sapVirtualInstanceName = "X00";
  const body = {
    location: "eastus",
    properties: {
      configuration: {
        appLocation: "eastus",
        configurationType: "Deployment",
        infrastructureConfiguration: {
          appResourceGroup: "X00-RG",
          databaseType: "HANA",
          deploymentType: "SingleServer",
          networkConfiguration: { isSecondaryIpEnabled: true },
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
      },
      environment: "NonProd",
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
