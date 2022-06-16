const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Generates a unique VPN profile for P2S clients for VirtualWan and associated VpnServerConfiguration combination in the specified resource group.
 *
 * @summary Generates a unique VPN profile for P2S clients for VirtualWan and associated VpnServerConfiguration combination in the specified resource group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/GenerateVirtualWanVpnServerConfigurationVpnProfile.json
 */
async function generateVirtualWanVpnServerConfigurationVpnProfile() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualWANName = "wan1";
  const vpnClientParams = {
    authenticationMethod: "EAPTLS",
    vpnServerConfigurationResourceId:
      "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/vpnServerConfigurations/vpnconfig1",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.beginGeneratevirtualwanvpnserverconfigurationvpnprofileAndWait(
    resourceGroupName,
    virtualWANName,
    vpnClientParams
  );
  console.log(result);
}

generateVirtualWanVpnServerConfigurationVpnProfile().catch(console.error);
