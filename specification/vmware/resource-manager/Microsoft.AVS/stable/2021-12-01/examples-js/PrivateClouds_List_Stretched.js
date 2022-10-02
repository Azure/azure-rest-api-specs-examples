const { AzureVMwareSolutionAPI } = require("@azure/arm-avs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List private clouds in a resource group
 *
 * @summary List private clouds in a resource group
 * x-ms-original-file: specification/vmware/resource-manager/Microsoft.AVS/stable/2021-12-01/examples/PrivateClouds_List_Stretched.json
 */
async function privateCloudsListStretched() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "group1";
  const credential = new DefaultAzureCredential();
  const client = new AzureVMwareSolutionAPI(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.privateClouds.list(resourceGroupName)) {
    resArray.push(item);
  }
  console.log(resArray);
}

privateCloudsListStretched().catch(console.error);
