const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates tags of a Security Partner Provider resource.
 *
 * @summary Updates tags of a Security Partner Provider resource.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/SecurityPartnerProviderUpdateTags.json
 */
async function updateSecurityPartnerProviderTags() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const securityPartnerProviderName = "securityPartnerProvider";
  const parameters = { tags: { tag1: "value1", tag2: "value2" } };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.securityPartnerProviders.updateTags(
    resourceGroupName,
    securityPartnerProviderName,
    parameters
  );
  console.log(result);
}

updateSecurityPartnerProviderTags().catch(console.error);
