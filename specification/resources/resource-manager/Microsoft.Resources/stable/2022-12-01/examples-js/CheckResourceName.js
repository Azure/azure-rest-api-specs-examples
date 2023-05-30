const { SubscriptionClient } = require("@azure/arm-resources-subscriptions");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to A resource name is valid if it is not a reserved word, does not contains a reserved word and does not start with a reserved word
 *
 * @summary A resource name is valid if it is not a reserved word, does not contains a reserved word and does not start with a reserved word
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2022-12-01/examples/CheckResourceName.json
 */
async function checkValidityForAResourceName() {
  const credential = new DefaultAzureCredential();
  const client = new SubscriptionClient(credential);
  const result = await client.checkResourceName();
  console.log(result);
}
