const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the specified Security Partner Provider.
 *
 * @summary Gets the specified Security Partner Provider.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/SecurityPartnerProviderGet.json
 */
async function getSecurityPartnerProvider() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const securityPartnerProviderName = "securityPartnerProvider";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.securityPartnerProviders.get(
    resourceGroupName,
    securityPartnerProviderName
  );
  console.log(result);
}

getSecurityPartnerProvider().catch(console.error);
