const { AzureReservationAPI } = require("@azure/arm-reservations");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a list of current quotas (service limits) and usage for all resources. The response from the list quota operation can be leveraged to request quota updates.
 *
 * @summary Gets a list of current quotas (service limits) and usage for all resources. The response from the list quota operation can be leveraged to request quota updates.
 * x-ms-original-file: specification/reservations/resource-manager/Microsoft.Capacity/stable/2020-10-25/examples/getMachineLearningServicesUsages.json
 */
async function quotasListUsagesMachineLearningServices() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const providerId = "Microsoft.MachineLearningServices";
  const location = "eastus";
  const credential = new DefaultAzureCredential();
  const client = new AzureReservationAPI(credential);
  const resArray = new Array();
  for await (let item of client.quota.list(subscriptionId, providerId, location)) {
    resArray.push(item);
  }
  console.log(resArray);
}
