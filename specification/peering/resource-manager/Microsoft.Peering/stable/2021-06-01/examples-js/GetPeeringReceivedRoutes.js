const { PeeringManagementClient } = require("@azure/arm-peering");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists the prefixes received over the specified peering under the given subscription and resource group.
 *
 * @summary Lists the prefixes received over the specified peering under the given subscription and resource group.
 * x-ms-original-file: specification/peering/resource-manager/Microsoft.Peering/stable/2021-06-01/examples/GetPeeringReceivedRoutes.json
 */
async function listsThePrefixesReceivedOverTheSpecifiedPeeringUnderTheGivenSubscriptionAndResourceGroup() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const peeringName = "peeringName";
  const prefix = "1.1.1.0/24";
  const asPath = "123 456";
  const originAsValidationState = "Valid";
  const rpkiValidationState = "Valid";
  const options = {
    prefix,
    asPath,
    originAsValidationState,
    rpkiValidationState,
  };
  const credential = new DefaultAzureCredential();
  const client = new PeeringManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.receivedRoutes.listByPeering(
    resourceGroupName,
    peeringName,
    options
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listsThePrefixesReceivedOverTheSpecifiedPeeringUnderTheGivenSubscriptionAndResourceGroup().catch(
  console.error
);
