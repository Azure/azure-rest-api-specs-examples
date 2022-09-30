const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the alert rule.
 *
 * @summary Creates or updates the alert rule.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateMicrosoftSecurityIncidentCreationAlertRule.json
 */
async function createsOrUpdatesAMicrosoftSecurityIncidentCreationRule() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const ruleId = "microsoftSecurityIncidentCreationRuleExample";
  const alertRule = {
    displayName: "testing displayname",
    enabled: true,
    etag: '"260097e0-0000-0d00-0000-5d6fa88f0000"',
    kind: "MicrosoftSecurityIncidentCreation",
    productFilter: "Microsoft Cloud App Security",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.alertRules.createOrUpdate(
    resourceGroupName,
    workspaceName,
    ruleId,
    alertRule
  );
  console.log(result);
}

createsOrUpdatesAMicrosoftSecurityIncidentCreationRule().catch(console.error);
