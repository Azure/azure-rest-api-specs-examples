const { WorkloadsClient } = require("@azure/arm-workloads");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates an Virtual Instance for SAP.
 *
 * @summary Creates an Virtual Instance for SAP.
 * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/sapvirtualinstances/SAPVirtualInstances_Create_SingleServer.json
 */
async function sapVirtualInstancesCreateSingleServer() {
  const subscriptionId = "8e17e36c-42e9-4cd5-a078-7b44883414e0";
  const resourceGroupName = "test-rg";
  const sapVirtualInstanceName = "X00";
  const body = {
    configuration: {
      appLocation: "eastus",
      configurationType: "Deployment",
      infrastructureConfiguration: {
        appResourceGroup: "X00-RG",
        databaseType: "HANA",
        deploymentType: "SingleServer",
        networkConfiguration: { isSecondaryIpEnabled: true },
        subnetId:
          "/subscriptions/49d64d54-e966-4c46-a868-1999802b762c/resourceGroups/dindurkhya-e2etesting/providers/Microsoft.Networks/virtualNetworks/test-vnet/subnets/appsubnet",
        virtualMachineConfiguration: {
          imageReference: {
            offer: "RHEL-SAP",
            publisher: "RedHat",
            sku: "7.4",
            version: "7.4.2019062505",
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
    location: "westcentralus",
    sapProduct: "S4HANA",
    tags: {},
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new WorkloadsClient(credential, subscriptionId);
  const result = await client.sAPVirtualInstances.beginCreateAndWait(
    resourceGroupName,
    sapVirtualInstanceName,
    options
  );
  console.log(result);
}

sapVirtualInstancesCreateSingleServer().catch(console.error);
