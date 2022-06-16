const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a record set within a DNS zone.
 *
 * @summary Creates or updates a record set within a DNS zone.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/CreateOrUpdateARecordsetAlias.json
 */
async function createARecordsetWithAliasTargetResource() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const zoneName = "zone1";
  const relativeRecordSetName = "record1";
  const recordType = "A";
  const parameters = {
    ttl: 3600,
    metadata: { key1: "value1" },
    targetResource: {
      id: "/subscriptions/726f8cd6-6459-4db4-8e6d-2cd2716904e2/resourceGroups/test/providers/Microsoft.Network/trafficManagerProfiles/testpp2",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.recordSets.createOrUpdate(
    resourceGroupName,
    zoneName,
    relativeRecordSetName,
    recordType,
    parameters
  );
  console.log(result);
}

createARecordsetWithAliasTargetResource().catch(console.error);
