const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update an metric alert definition.
 *
 * @summary Create or update an metric alert definition.
 * x-ms-original-file: specification/monitor/resource-manager/Microsoft.Insights/stable/2018-03-01/examples/createOrUpdateWebTestMetricAlert.json
 */
async function createOrUpdateAWebTestAlertRule() {
  const subscriptionId =
    process.env["MONITOR_SUBSCRIPTION_ID"] || "12345678-1234-1234-1234-123456789101";
  const resourceGroupName = process.env["MONITOR_RESOURCE_GROUP"] || "rg-example";
  const ruleName = "webtest-name-example";
  const parameters = {
    description: 'Automatically created alert rule for availability test "component-example" a',
    actions: [],
    criteria: {
      componentId:
        "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example",
      failedLocationCount: 2,
      odataType: "Microsoft.Azure.Monitor.WebtestLocationAvailabilityCriteria",
      webTestId:
        "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example",
    },
    enabled: true,
    evaluationFrequency: "PT1M",
    location: "global",
    scopes: [
      "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example",
      "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example",
    ],
    severity: 4,
    tags: {
      "hiddenLink:/subscriptions/12345678123412341234123456789101/resourcegroups/rgExample/providers/microsoftInsights/components/webtestNameExample":
        "Resource",
      "hiddenLink:/subscriptions/12345678123412341234123456789101/resourcegroups/rgExample/providers/microsoftInsights/webtests/componentExample":
        "Resource",
    },
    windowSize: "PT15M",
  };
  const credential = new DefaultAzureCredential();
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.createOrUpdate(resourceGroupName, ruleName, parameters);
  console.log(result);
}
