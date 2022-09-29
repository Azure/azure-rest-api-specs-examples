const { SecurityInsights } = require("@azure/arm-securityinsight");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates the alert rule.
 *
 * @summary Creates or updates the alert rule.
 * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-09-01-preview/examples/alertRules/CreateFusionAlertRuleWithFusionScenarioExclusion.json
 */
async function createsOrUpdatesAFusionAlertRuleWithScenarioExclusionPattern() {
  const subscriptionId = "d0cfe6b2-9ac0-4464-9919-dccaee2e48c0";
  const resourceGroupName = "myRg";
  const workspaceName = "myWorkspace";
  const ruleId = "myFirstFusionRule";
  const alertRule = {
    alertRuleTemplateName: "f71aba3d-28fb-450b-b192-4e76a83015c8",
    enabled: true,
    etag: "3d00c3ca-0000-0100-0000-5d42d5010000",
    kind: "Fusion",
    sourceSettings: [
      { enabled: true, sourceName: "Anomalies", sourceSubTypes: [] },
      {
        enabled: true,
        sourceName: "Alert providers",
        sourceSubTypes: [
          {
            enabled: true,
            severityFilters: {
              filters: [
                { enabled: true, severity: "High" },
                { enabled: true, severity: "Medium" },
                { enabled: true, severity: "Low" },
                { enabled: true, severity: "Informational" },
              ],
            },
            sourceSubTypeName: "Azure Active Directory Identity Protection",
          },
          {
            enabled: true,
            severityFilters: {
              filters: [
                { enabled: true, severity: "High" },
                { enabled: true, severity: "Medium" },
                { enabled: true, severity: "Low" },
                { enabled: true, severity: "Informational" },
              ],
            },
            sourceSubTypeName: "Azure Defender",
          },
          {
            enabled: true,
            severityFilters: {
              filters: [
                { enabled: true, severity: "High" },
                { enabled: true, severity: "Medium" },
                { enabled: true, severity: "Low" },
                { enabled: true, severity: "Informational" },
              ],
            },
            sourceSubTypeName: "Azure Defender for IoT",
          },
          {
            enabled: true,
            severityFilters: {
              filters: [
                { enabled: true, severity: "High" },
                { enabled: true, severity: "Medium" },
                { enabled: true, severity: "Low" },
                { enabled: true, severity: "Informational" },
              ],
            },
            sourceSubTypeName: "Microsoft 365 Defender",
          },
          {
            enabled: true,
            severityFilters: {
              filters: [
                { enabled: true, severity: "High" },
                { enabled: true, severity: "Medium" },
                { enabled: true, severity: "Low" },
                { enabled: true, severity: "Informational" },
              ],
            },
            sourceSubTypeName: "Microsoft Cloud App Security",
          },
          {
            enabled: true,
            severityFilters: {
              filters: [
                { enabled: true, severity: "High" },
                { enabled: true, severity: "Medium" },
                { enabled: true, severity: "Low" },
                { enabled: true, severity: "Informational" },
              ],
            },
            sourceSubTypeName: "Microsoft Defender for Endpoint",
          },
          {
            enabled: true,
            severityFilters: {
              filters: [
                { enabled: true, severity: "High" },
                { enabled: true, severity: "Medium" },
                { enabled: true, severity: "Low" },
                { enabled: true, severity: "Informational" },
              ],
            },
            sourceSubTypeName: "Microsoft Defender for Identity",
          },
          {
            enabled: true,
            severityFilters: {
              filters: [
                { enabled: true, severity: "High" },
                { enabled: true, severity: "Medium" },
                { enabled: true, severity: "Low" },
                { enabled: true, severity: "Informational" },
              ],
            },
            sourceSubTypeName: "Microsoft Defender for Office 365",
          },
          {
            enabled: true,
            severityFilters: {
              filters: [
                { enabled: true, severity: "High" },
                { enabled: true, severity: "Medium" },
                { enabled: true, severity: "Low" },
                { enabled: true, severity: "Informational" },
              ],
            },
            sourceSubTypeName: "Azure Sentinel scheduled analytics rules",
          },
        ],
      },
      {
        enabled: true,
        sourceName: "Raw logs from other sources",
        sourceSubTypes: [
          {
            enabled: true,
            severityFilters: { filters: [] },
            sourceSubTypeName: "Palo Alto Networks",
          },
        ],
      },
    ],
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

createsOrUpdatesAFusionAlertRuleWithScenarioExclusionPattern().catch(console.error);
