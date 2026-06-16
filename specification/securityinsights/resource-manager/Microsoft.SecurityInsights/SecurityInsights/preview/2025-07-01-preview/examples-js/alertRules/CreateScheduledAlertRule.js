const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates the alert rule.
 *
 * @summary creates or updates the alert rule.
 * x-ms-original-file: 2025-07-01-preview/alertRules/CreateScheduledAlertRule.json
 */
async function createsOrUpdatesAScheduledAlertRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const client = new SecurityInsights(credential, subscriptionId);
  const result = await client.alertRules.createOrUpdate(
    "myRg",
    "myWorkspace",
    "73e01a99-5cd7-4139-a149-9f2736ff2ab5",
    {
      etag: '"0300bf09-0000-0000-0000-5c37296e0000"',
      kind: "Scheduled",
      description: "An example for a scheduled rule",
      alertDetailsOverride: {
        alertDescriptionFormat: "Suspicious activity was made by {{ComputerIP}}",
        alertDisplayNameFormat: "Alert from {{Computer}}",
        alertDynamicProperties: [
          { alertProperty: "ProductComponentName", value: "ProductComponentNameCustomColumn" },
          { alertProperty: "ProductName", value: "ProductNameCustomColumn" },
          { alertProperty: "AlertLink", value: "Link" },
        ],
      },
      customDetails: { OperatingSystemName: "OSName", OperatingSystemType: "OSType" },
      displayName: "My scheduled rule",
      enabled: true,
      entityMappings: [
        { entityType: "Host", fieldMappings: [{ columnName: "Computer", identifier: "FullName" }] },
        { entityType: "IP", fieldMappings: [{ columnName: "ComputerIP", identifier: "Address" }] },
      ],
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
      query: "Heartbeat",
      queryFrequency: "PT1H",
      queryPeriod: "P2DT1H30M",
      sentinelEntitiesMappings: [{ columnName: "Entities" }],
      severity: "High",
      suppressionDuration: "PT1H",
      suppressionEnabled: false,
      tactics: ["Persistence", "LateralMovement"],
      techniques: ["T1037", "T1021"],
      triggerOperator: "GreaterThan",
      triggerThreshold: 0,
    },
  );
  console.log(result);
}
