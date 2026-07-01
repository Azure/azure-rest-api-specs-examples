const { MonitorClient } = require("@azure/arm-monitor");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to create or update an metric alert definition.
 *
 * @summary create or update an metric alert definition.
 * x-ms-original-file: 2024-03-01-preview/createOrUpdateWebTestMetricAlert.json
 */
async function createOrUpdateAWebTestAlertRule() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "12345678-1234-1234-1234-123456789101";
  const client = new MonitorClient(credential, subscriptionId);
  const result = await client.metricAlerts.createOrUpdate("rg-example", "webtest-name-example", {
    location: "global",
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
    scopes: [
      "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example",
      "/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example",
    ],
    severity: 4,
    windowSize: "PT15M",
    tags: {
      "hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/components/webtest-name-example":
        "Resource",
      "hidden-link:/subscriptions/12345678-1234-1234-1234-123456789101/resourcegroups/rg-example/providers/microsoft.insights/webtests/component-example":
        "Resource",
    },
  });
  console.log(result);
}
