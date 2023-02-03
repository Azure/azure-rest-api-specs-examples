const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the alert rule.
 *
 * @summary Creates or updates the alert rule.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateScheduledAlertRule.json
 */
async function createsOrUpdatesAScheduledAlertRule() {
  const subscriptionId =
    process.env["SECURITYINSIGHT_SUBSCRIPTION_ID"] || "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = process.env["SECURITYINSIGHT_RESOURCE_GROUP"] || "myRg";
  const workspaceName = "myWorkspace";
  const ruleId = "73e01a99-5cd7-4139-a149-9f2736ff2ab5";
  const alertRule = {
    description: "An example for a scheduled rule",
    alertDetailsOverride: {
      alertDescriptionFormat: "Suspicious activity was made by {{ComputerIP}}",
      alertDisplayNameFormat: "Alert from {{Computer}}",
    },
    customDetails: {
      operatingSystemName: "OSName",
      operatingSystemType: "OSType",
    },
    displayName: "My scheduled rule",
    enabled: true,
    entityMappings: [
      {
        entityType: "Host",
        fieldMappings: [{ columnName: "Computer", identifier: "FullName" }],
      },
      {
        entityType: "IP",
        fieldMappings: [{ columnName: "ComputerIP", identifier: "Address" }],
      },
    ],
    etag: '"0300bf09-0000-0000-0000-5c37296e0000"',
    eventGroupingSettings: { aggregationKind: "AlertPerResult" },
    incidentConfiguration: {
      createIncident: true,
      groupingConfiguration: {
        enabled: true,
        groupByAlertDetails: ["DisplayName"],
        groupByCustomDetails: ["OperatingSystemType", "OperatingSystemName"],
        groupByEntities: ["Host"],
        lookbackDuration: "PT5H",
        matchingMethod: "Selected",
        reopenClosedIncident: false,
      },
    },
    kind: "Scheduled",
    query: "Heartbeat",
    queryFrequency: "PT1H",
    queryPeriod: "P2DT1H30M",
    severity: "High",
    suppressionDuration: "PT1H",
    suppressionEnabled: false,
    tactics: ["Persistence", "LateralMovement"],
    techniques: ["T1037", "T1021"],
    triggerOperator: "GreaterThan",
    triggerThreshold: 0,
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
