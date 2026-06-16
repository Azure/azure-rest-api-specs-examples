const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the alert rule.
 *
 * @summary creates or updates the alert rule.
 * x-ms-original-file: 2025-07-01-preview/alertRules/CreateNrtAlertRule.json
 */
async function createsOrUpdatesANrtAlertRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.alertRules.createOrUpdate(
    "myRg",
    "myWorkspace",
    "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
    {
      etag: '"0300bf09-0000-0000-0000-5c37296e0000"',
      kind: "NRT",
      description: "",
      displayName: "Rule2",
      enabled: true,
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
      query:
        "ProtectionStatus | extend HostCustomEntity = Computer | extend IPCustomEntity = ComputerIP_Hidden",
      severity: "High",
      suppressionDuration: "PT1H",
      suppressionEnabled: false,
      tactics: ["Persistence", "LateralMovement"],
      techniques: ["T1037", "T1021"],
    },
  );
  console.log(result);
}
