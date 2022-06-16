const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified private dns zone group.
 *
 * @summary Deletes the specified private dns zone group.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-05-01/examples/PrivateEndpointDnsZoneGroupDelete.json
 */
async function deletePrivateDnsZoneGroup() {
  const subscriptionId = "subId";
  const resourceGroupName = "rg1";
  const privateEndpointName = "testPe";
  const privateDnsZoneGroupName = "testPdnsgroup";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateDnsZoneGroups.beginDeleteAndWait(
    resourceGroupName,
    privateEndpointName,
    privateDnsZoneGroupName
  );
  console.log(result);
}

deletePrivateDnsZoneGroup().catch(console.error);
