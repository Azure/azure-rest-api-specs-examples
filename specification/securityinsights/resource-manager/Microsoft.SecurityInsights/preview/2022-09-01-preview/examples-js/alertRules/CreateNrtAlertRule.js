const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the alert rule.
 *
 * @summary Creates or updates the alert rule.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateNrtAlertRule.json
 */
async function createsOrUpdatesANrtAlertRule() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const ruleId = "73e01a99-5cd7-4139-a149-9f2736ff2ab5";
  const alertRule = {
    description: "",
    displayName: "Rule2",
    enabled: true,
    etag: '"0300bf09-0000-0000-0000-5c37296e0000"',
    eventGroupingSettings: { aggregationKind: "AlertPerResult" },
    incidentConfiguration: {
      createIncident: true,
      groupingConfiguration: {
        enabled: true,
        groupByEntities: ["Host", "Account"],
        lookbackDuration: "PT5H",
        matchingMethod: "Selected",
        reopenClosedIncident: false,
      },
    },
    kind: "NRT",
    query:
      "ProtectionStatus | extend HostCustomEntity = Computer | extend IPCustomEntity = ComputerIP_Hidden",
    severity: "High",
    suppressionDuration: "PT1H",
    suppressionEnabled: false,
    tactics: ["Persistence", "LateralMovement"],
    techniques: ["T1037", "T1021"],
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
