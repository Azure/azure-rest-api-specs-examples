const createNetworkManagementClient = require("@azure-rest/arm-network").default,
  { paginate } = require("@azure-rest/arm-network");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Gets all Network Virtual Appliances in a subscription.
 *
 * @summary Gets all Network Virtual Appliances in a subscription.
 * x-ms-original-file: specification/network/resource-manager/Microsoft.Network/stable/2022-05-01/examples/NetworkVirtualApplianceListBySubscription.json
 */
async function listAllNetworkVirtualAppliancesForAGivenSubscription() {
  const credential = new DefaultAzureCredential();
  const client = createNetworkManagementClient(credential);
  const subscriptionId = "";
  const options = {
    queryParameters: { "api-version": "2022-05-01" },
  };
  const initialResponse = await client
    .path(
      "/subscriptions/{subscriptionId}/providers/Microsoft.Network/networkVirtualAppliances",
      subscriptionId,
    )
    .get(options);
  const pageData = paginate(client, initialResponse);
  const result = [];
  for await (const item of pageData) {
    result.push(item);
  }
  console.log(result);
}

listAllNetworkVirtualAppliancesForAGivenSubscription().catch(console.error);
