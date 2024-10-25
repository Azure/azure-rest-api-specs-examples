const { DnsManagementClient } = require("@azure/arm-dns");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns the DNS records specified by the referencing targetResourceIds.
 *
 * @summary Returns the DNS records specified by the referencing targetResourceIds.
 * x-ms-original-file: specification/dns/resource-manager/Microsoft.Network/preview/2023-07-01-preview/examples/GetDnsResourceReference.json
 */
async function getDnsResourceReference() {
  const subscriptionId = process.env["DNS_SUBSCRIPTION_ID"] || "subid";
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
