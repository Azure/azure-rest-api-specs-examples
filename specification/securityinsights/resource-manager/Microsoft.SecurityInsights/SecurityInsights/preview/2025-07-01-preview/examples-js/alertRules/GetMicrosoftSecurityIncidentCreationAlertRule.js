const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to gets the alert rule.
 *
 * @summary gets the alert rule.
 * x-ms-original-file: 2025-07-01-preview/alertRules/GetMicrosoftSecurityIncidentCreationAlertRule.json
 */
async function getAMicrosoftSecurityIncidentCreationRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.alertRules.get(
    "myRg",
    "myWorkspace",
    "microsoftSecurityIncidentCreationRuleExample",
  );
  console.log(result);
}
