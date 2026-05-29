const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to use this method to get the list of IoT Security solutions by subscription.
 *
 * @summary use this method to get the list of IoT Security solutions by subscription.
 * x-ms-original-file: 2019-08-01/IoTSecuritySolutions/GetIoTSecuritySolutionsList.json
 */
async function listIoTSecuritySolutionsBySubscription() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (const item of client.iotSecuritySolution.listBySubscription()) {
    resArray.push(item);
  }

  console.log(resArray);
}
