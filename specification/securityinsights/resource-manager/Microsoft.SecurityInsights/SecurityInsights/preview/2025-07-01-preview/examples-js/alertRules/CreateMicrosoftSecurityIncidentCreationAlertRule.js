const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the alert rule.
 *
 * @summary creates or updates the alert rule.
 * x-ms-original-file: 2025-07-01-preview/alertRules/CreateMicrosoftSecurityIncidentCreationAlertRule.json
 */
async function createsOrUpdatesAMicrosoftSecurityIncidentCreationRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.alertRules.createOrUpdate(
    "myRg",
    "myWorkspace",
    "microsoftSecurityIncidentCreationRuleExample",
    {
      etag: '"260097e0-0000-0d00-0000-5d6fa88f0000"',
      kind: "MicrosoftSecurityIncidentCreation",
      displayName: "testing displayname",
      enabled: true,
      productFilter: "Microsoft Cloud App Security",
    },
  );
  console.log(result);
}
