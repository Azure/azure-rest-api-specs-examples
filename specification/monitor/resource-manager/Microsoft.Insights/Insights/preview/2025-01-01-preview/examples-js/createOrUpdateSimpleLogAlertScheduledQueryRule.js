const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to creates or updates a scheduled query rule.
 *
 * @summary creates or updates a scheduled query rule.
 * x-ms-original-file: 2025-01-01-preview/createOrUpdateSimpleLogAlertScheduledQueryRule.json
 */
async function createOrUpdateASimpleLogAlertScheduledQueryRuleOnSubscription() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "dd4bfc94-a096-412b-9c43-4bd13e35afbc";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.scheduledQueryRules.createOrUpdate("QueryResourceGroupName", "perf", {
    kind: "SimpleLogAlert",
    location: "eastus",
    description: "Performance rule",
    actions: {
      actionGroups: [
        "/subscriptions/1cf177ed-1330-4692-80ea-fd3d7783b147/resourcegroups/sqrapi/providers/microsoft.insights/actiongroups/myactiongroup",
      ],
      actionProperties: { "Icm.Title": "Custom title in ICM", "Icm.TsgId": "https://tsg.url" },
      customProperties: { key11: "value11", key12: "value12" },
    },
    autoMitigate: false,
    checkWorkspaceAlertsStorageConfigured: true,
    criteria: { allOf: [{ query: 'Perf | where ObjectName == "Processor"' }] },
    enabled: true,
    scopes: [
      "/subscriptions/aaf177ed-1330-a9f2-80ea-fd3d7783b147/resourceGroups/scopeResourceGroup1/providers/Microsoft.Compute/virtualMachines/vm1",
    ],
    severity: 4,
    skipQueryValidation: true,
  });
  console.log(result);
}
