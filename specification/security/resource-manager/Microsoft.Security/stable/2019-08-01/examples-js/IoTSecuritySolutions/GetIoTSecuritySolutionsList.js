const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Use this method to get the list of IoT Security solutions by subscription.
 *
 * @summary Use this method to get the list of IoT Security solutions by subscription.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2019-08-01/examples/IoTSecuritySolutions/GetIoTSecuritySolutionsList.json
 */
async function listIoTSecuritySolutionsBySubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.iotSecuritySolution.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listIoTSecuritySolutionsBySubscription().catch(console.error);
