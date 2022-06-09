```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update existing rule or create new rule if it doesn't exist
 *
 * @summary Update existing rule or create new rule if it doesn't exist
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2019-01-01-preview/examples/AlertsSuppressionRules/PutAlertsSuppressionRule_example.json
 */
async function updateOrCreateSuppressionRuleForSubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const alertsSuppressionRuleName = "dismissIpAnomalyAlerts";
  const alertsSuppressionRule = {
    alertType: "IpAnomaly",
    comment: "Test VM",
    expirationDateUtc: new Date("2019-12-01T19:50:47.083633Z"),
    reason: "FalsePositive",
    state: "Enabled",
    suppressionAlertsScope: {
      allOf: [{ field: "entities.ip.address" }, { field: "entities.process.commandline" }],
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.alertsSuppressionRules.update(
    alertsSuppressionRuleName,
    alertsSuppressionRule
  );
  console.log(result);
}

updateOrCreateSuppressionRuleForSubscription().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.
