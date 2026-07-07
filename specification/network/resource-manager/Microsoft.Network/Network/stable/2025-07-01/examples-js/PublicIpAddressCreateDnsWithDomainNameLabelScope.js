const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a static or dynamic public IP address.
 *
 * @summary creates or updates a static or dynamic public IP address.
 * x-ms-original-file: 2025-07-01/PublicIpAddressCreateDnsWithDomainNameLabelScope.json
 */
async function createPublicIPAddressDNSWithDomainNameLabelScope() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.publicIPAddresses.createOrUpdate("rg1", "test-ip", {
    location: "eastus",
    dnsSettings: { domainNameLabel: "dnslbl", domainNameLabelScope: "TenantReuse" },
  });
  console.log(result);
}
