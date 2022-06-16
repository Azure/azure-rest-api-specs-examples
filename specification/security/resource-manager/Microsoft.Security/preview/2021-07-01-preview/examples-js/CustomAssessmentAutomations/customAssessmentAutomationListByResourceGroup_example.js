const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to List custom assessment automations by provided subscription and resource group
 *
 * @summary List custom assessment automations by provided subscription and resource group
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomAssessmentAutomations/customAssessmentAutomationListByResourceGroup_example.json
 */
async function listCustomAssessmentAutomationsInASubscriptionAndAResourceGroup() {
  const subscriptionId = "e5d1b86c-3051-44d5-8802-aa65d45a279b";
  const resourceGroupName = "TestResourceGroup";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.customAssessmentAutomations.listByResourceGroup(
    resourceGroupName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

listCustomAssessmentAutomationsInASubscriptionAndAResourceGroup().catch(console.error);
