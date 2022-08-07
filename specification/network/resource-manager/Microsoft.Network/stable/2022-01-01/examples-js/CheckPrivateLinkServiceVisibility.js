const { NetworkManagementClient } = require("@azure/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Checks whether the subscription is visible to private link service.
 *
 * @summary Checks whether the subscription is visible to private link service.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-01-01/examples/CheckPrivateLinkServiceVisibility.json
 */
async function checkPrivateLinkServiceVisibility() {
  const subscriptionId = "subid";
  const location = "westus";
  const parameters = {
    privateLinkServiceAlias: "mypls.00000000-0000-0000-0000-000000000000.azure.privatelinkservice",
  };
  const credential = new DefaultAzureCredential();
  const client = new NetworkManagementClient(credential, subscriptionId);
  const result = await client.privateLinkServices.beginCheckPrivateLinkServiceVisibilityAndWait(
    location,
    parameters
  );
  console.log(result);
}

checkPrivateLinkServiceVisibility().catch(console.error);
