const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a security automation.
 *
 * @summary Deletes a security automation.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/Automations/DeleteAutomation_example.json
 */
async function deleteASecurityAutomation() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const resourceGroupName = "myRg";
  const automationName = "myAutomationName";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.automations.delete(resourceGroupName, automationName);
  console.log(result);
}

deleteASecurityAutomation().catch(console.error);
