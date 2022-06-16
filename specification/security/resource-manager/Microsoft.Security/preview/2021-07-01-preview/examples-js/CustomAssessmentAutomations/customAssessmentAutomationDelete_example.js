const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a custom assessment automation by name for a provided subscription
 *
 * @summary Deletes a custom assessment automation by name for a provided subscription
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomAssessmentAutomations/customAssessmentAutomationDelete_example.json
 */
async function deleteACustomAssessmentAutomation() {
  const subscriptionId = "e5d1b86c-3051-44d5-8802-aa65d45a279b";
  const resourceGroupName = "TestResourceGroup";
  const customAssessmentAutomationName = "MyCustomAssessmentAutomation";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.customAssessmentAutomations.delete(
    resourceGroupName,
    customAssessmentAutomationName
  );
  console.log(result);
}

deleteACustomAssessmentAutomation().catch(console.error);
