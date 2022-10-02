const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List private clouds in a subscription
 *
 * @summary List private clouds in a subscription
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_ListInSubscription_Stretched.json
 */
async function privateCloudsListInSubscriptionStretched() {
  const subscriptionId = "{subscription-id}";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateClouds.listInSubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

privateCloudsListInSubscriptionStretched().catch(console.error);
