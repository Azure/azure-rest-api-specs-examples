Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-network_27.0.0/sdk/network/arm-network/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to The Set VpnclientIpsecParameters operation sets the vpnclient ipsec policy for P2S client of virtual network gateway in the specified resource group through Network resource provider.
 *
 * @summary The Set VpnclientIpsecParameters operation sets the vpnclient ipsec policy for P2S client of virtual network gateway in the specified resource group through Network resource provider.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/VirtualNetworkGatewaySetVpnClientIpsecParameters.json
 */
async function setVirtualNetworkGatewayVpnClientIpsecParameters() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const virtualNetworkGatewayName = "vpngw";
  const vpnclientIpsecParams = {
    dhGroup: "DHGroup2",
    ikeEncryption: "AES256",
    ikeIntegrity: "SHA384",
    ipsecEncryption: "AES256",
    ipsecIntegrity: "SHA256",
    pfsGroup: "PFS2",
    saDataSizeKilobytes: 429497,
    saLifeTimeSeconds: 86473,
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkGateways.beginSetVpnclientIpsecParametersAndWait(
    resourceGroupName,
    virtualNetworkGatewayName,
    vpnclientIpsecParams
  );
  console.log(result);
}

setVirtualNetworkGatewayVpnClientIpsecParameters().catch(console.error);
```
