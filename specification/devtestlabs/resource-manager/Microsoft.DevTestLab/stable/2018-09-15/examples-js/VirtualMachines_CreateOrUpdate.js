const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace an existing virtual machine. This operation can take a while to complete.
 *
 * @summary Create or replace an existing virtual machine. This operation can take a while to complete.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/VirtualMachines_CreateOrUpdate.json
 */
async function virtualMachinesCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const name = "{vmName}";
  const labVirtualMachine = {
    allowClaim: true,
    disallowPublicIpAddress: true,
    galleryImageReference: {
      offer: "UbuntuServer",
      osType: "Linux",
      publisher: "Canonical",
      sku: "16.04-LTS",
      version: "Latest",
    },
    labSubnetName: "{virtualNetworkName}Subnet",
    labVirtualNetworkId:
      "/subscriptions/{subscriptionId}/resourcegroups/resourceGroupName/providers/microsoft.devtestlab/labs/{labName}/virtualnetworks/{virtualNetworkName}",
    location: "{location}",
    password: "{userPassword}",
    size: "Standard_A2_v2",
    storageType: "Standard",
    tags: { tagName1: "tagValue1" },
    userName: "{userName}",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.virtualMachines.beginCreateOrUpdateAndWait(
    resourceGroupName,
    labName,
    name,
    labVirtualMachine
  );
  console.log(result);
}

virtualMachinesCreateOrUpdate().catch(console.error);
