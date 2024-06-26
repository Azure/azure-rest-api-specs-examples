const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve an alert rule status.
 *
 * @summary Retrieve an alert rule status.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/getMetricAlertStatusByName.json
 */
async function getAnAlertRuleStatus() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "009f6022-67ec-423e-9aa7-691182870588";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "EastUs";
  const ruleName = "custom1";
  const statusName =
    "cmVzb3VyY2VJZD0vc3Vic2NyaXB0aW9ucy8xNGRkZjBjNS03N2M1LTRiNTMtODRmNi1lMWZhNDNhZDY4ZjcvcmVzb3VyY2VHcm91cHMvZ2lndGVzdC9wcm92aWRlcnMvTWljcm9zb2Z0LkNvbXB1dGUvdmlydHVhbE1hY2hpbmVzL2dpZ3dhZG1l";
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlertsStatus.listByName(
    resourceGroupName,
    ruleName,
    statusName,
  );
  console.log(result);
}
