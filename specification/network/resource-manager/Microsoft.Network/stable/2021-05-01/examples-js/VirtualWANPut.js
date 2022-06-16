const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a VirtualWAN resource if it doesn't exist else updates the existing VirtualWAN.
 *
 * @summary Creates a VirtualWAN resource if it doesn't exist else updates the existing VirtualWAN.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualWANPut.json
 */
async function virtualWanCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualWANName = "wan1";
  const wANParameters = {
    typePropertiesType: "Basic",
    disableVpnEncryption: false,
    location: "West US",
    tags: { key1: "value1" },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualWans.beginCreateOrUpdateAndWait(
    resourceGroupName,
    virtualWANName,
    wANParameters
  );
  console.log(result);
}

virtualWanCreate().catch(console.error);
