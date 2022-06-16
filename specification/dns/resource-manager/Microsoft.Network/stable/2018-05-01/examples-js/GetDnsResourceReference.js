const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the DNS records specified by the referencing targetResourceIds.
 *
 * @summary Returns the DNS records specified by the referencing targetResourceIds.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/stable/2018-05-01/examples/GetDnsResourceReference.json
 */
async function listZonesByResourceGroup() {
  const subscriptionId = "subid";
  const parameters = {
    targetResources: [
      {
        id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/trafficManagerProfiles/testpp2",
      },
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new DnsManagementClient(credential, subscriptionId);
  const result = await client.dnsResourceReferenceOperations.getByTargetResources(parameters);
  console.log(result);
}

listZonesByResourceGroup().catch(console.error);
