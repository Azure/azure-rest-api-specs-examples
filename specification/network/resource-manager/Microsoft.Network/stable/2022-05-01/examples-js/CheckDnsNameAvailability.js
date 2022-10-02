const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks whether a domain name in the cloudapp.azure.com zone is available for use.
 *
 * @summary Checks whether a domain name in the cloudapp.azure.com zone is available for use.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/CheckDnsNameAvailability.json
 */
async function checkDnsNameAvailability() {
  const subscriptionId = "subid";
  const location = "westus";
  const domainNameLabel = "testdns";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.checkDnsNameAvailability(location, domainNameLabel);
  console.log(result);
}

checkDnsNameAvailability().catch(console.error);
