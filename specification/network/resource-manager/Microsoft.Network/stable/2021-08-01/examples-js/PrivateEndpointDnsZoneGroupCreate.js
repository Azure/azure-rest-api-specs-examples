const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a private dns zone group in the specified private endpoint.
 *
 * @summary Creates or updates a private dns zone group in the specified private endpoint.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/PrivateEndpointDnsZoneGroupCreate.json
 */
async function createPrivateDnsZoneGroup() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const privateEndpointName = "testPe";
  const privateDnsZoneGroupName = "testPdnsgroup";
  const parameters = {
    privateDnsZoneConfigs: [
      {
        privateDnsZoneId:
          "/subscriptions/subId/resourceGroups/rg1/providers/Microsoft.Network/privateDnsZones/zone1.com",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateDnsZoneGroups.beginCreateOrUpdateAndWait(
    resourceGroupName,
    privateEndpointName,
    privateDnsZoneGroupName,
    parameters
  );
  console.log(result);
}

createPrivateDnsZoneGroup().catch(console.error);
