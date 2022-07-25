const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Delete the office365 consent.
 *
 * @summary Delete the office365 consent.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-07-01-preview/examples/officeConsents/DeleteOfficeConsents.json
 */
async function deleteAnOfficeConsent() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const consentId = "04e5fd05-ff86-4b97-b8d2-1c20933cb46c";
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.officeConsents.delete(resourceGroupName, workspaceName, consentId);
  console.log(result);
}

deleteAnOfficeConsent().catch(console.error);
