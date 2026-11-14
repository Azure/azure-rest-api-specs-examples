const { PrivateDnsManagementClient } = require("@azure/arm-privatedns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a record set within a Private DNS zone.
 *
 * @summary creates or updates a record set within a Private DNS zone.
 * x-ms-original-file: 2026-06-01-preview/PTMLinkRecordSetAPut.json
 */
async function putPrivateDNSZonePTMLinkARecordSet() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "subscriptionId";
  const client = new PrivateDnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
    "resourceGroup1",
    "privatezone1.com",
    "A",
    "ptmlinkrecord",
    {
      metadata: { key1: "value1" },
      ttl: 3600,
      trafficManagementProfile: {
        id: "/subscriptions/subscriptionId/resourceGroups/resourceGroup1/providers/Microsoft.Network/privateTrafficManagerProfiles/ptm-profile1",
      },
    },
  );
  console.log(result);
}
