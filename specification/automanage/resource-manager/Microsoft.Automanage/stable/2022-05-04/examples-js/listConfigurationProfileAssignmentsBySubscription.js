const { AutomanageClient } = require("@azure/arm-automanage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get list of configuration profile assignments under a given subscription
 *
 * @summary Get list of configuration profile assignments under a given subscription
 * x-ms-original-file: specification/automanage/resource-manager/Microsoft.Automanage/stable/2022-05-04/examples/listConfigurationProfileAssignmentsBySubscription.json
 */
async function listConfigurationProfileAssignmentsBySubscription() {
  const subscriptionId = "mySubscriptionId";
  const credential = new DefaultAzureCredential();
  const client = new AutomanageClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.configurationProfileAssignments.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listConfigurationProfileAssignmentsBySubscription().catch(console.error);
