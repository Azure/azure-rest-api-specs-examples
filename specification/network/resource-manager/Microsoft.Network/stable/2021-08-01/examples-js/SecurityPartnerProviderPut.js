const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the specified Security Partner Provider.
 *
 * @summary Creates or updates the specified Security Partner Provider.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2021-08-01/examples/SecurityPartnerProviderPut.json
 */
async function createSecurityPartnerProvider() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const securityPartnerProviderName = "securityPartnerProvider";
  const parameters = {
    location: "West US",
    securityProviderName: "ZScaler",
    tags: { key1: "value1" },
    virtualHub: {
      id: "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.Network/virtualHubs/hub1",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.securityPartnerProviders.beginCreateOrUpdateAndWait(
    resourceGroupName,
    securityPartnerProviderName,
    parameters
  );
  console.log(result);
}

createSecurityPartnerProvider().catch(console.error);
