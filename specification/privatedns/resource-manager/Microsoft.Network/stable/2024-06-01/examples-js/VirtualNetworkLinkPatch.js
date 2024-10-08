const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates a virtual network link to the specified Private DNS zone.
 *
 * @summary Updates a virtual network link to the specified Private DNS zone.
 * x-ms-original-file: specification/privatedns/resource-manager/Microsoft.Network/stable/2024-06-01/examples/VirtualNetworkLinkPatch.json
 */
async function patchPrivateDnsZoneVirtualNetworkLink() {
  const subscriptionId = process.env["PRIVATEDNS_SUBSCRIPTION_ID"] || "subscriptionId";
  const resourceGroupName = process.env["PRIVATEDNS_RESOURCE_GROUP"] || "resourceGroup1";
  const privateZoneName = "privatelink.contoso.com";
  const virtualNetworkLinkName = "virtualNetworkLink1";
  const parameters = {
    registrationEnabled: true,
    resolutionPolicy: "NxDomainRedirect",
    tags: { key2: "value2" },
  };
  const credential = new DefaultAzureCredential();
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.virtualNetworkLinks.beginUpdateAndWait(
    resourceGroupName,
    privateZoneName,
    virtualNetworkLinkName,
    parameters,
  );
  console.log(result);
}
