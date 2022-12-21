const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the specified Security Partner Provider.
 *
 * @summary Deletes the specified Security Partner Provider.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-07-01/examples/SecurityPartnerProviderDelete.json
 */
async function deleteSecurityPartnerProvider() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const securityPartnerProviderName = "securityPartnerProvider";
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.securityPartnerProviders.beginDeleteAndWait(
    resourceGroupName,
    securityPartnerProviderName
  );
  console.log(result);
}

deleteSecurityPartnerProvider().catch(console.error);
