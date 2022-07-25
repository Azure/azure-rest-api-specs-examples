const { LabServicesClient } = require("@azure/arm-labservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns a list of all lab plans within a subscription
 *
 * @summary Returns a list of all lab plans within a subscription
 * x-ms-original-file: specification/labservices/resource-manager/Microsoft.LabServices/preview/2021-11-15-preview/examples/LabPlans/listLabPlans.json
 */
async function listLabPlans() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const credential = new DefaultAzureCredential();
  const client = new LabServicesClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.labPlans.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listLabPlans().catch(console.error);
