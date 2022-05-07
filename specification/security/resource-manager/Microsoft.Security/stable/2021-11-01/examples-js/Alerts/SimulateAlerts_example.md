Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Simulate security alerts
 *
 * @summary Simulate security alerts
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-11-01/examples/Alerts/SimulateAlerts_example.json
 */
async function simulateSecurityAlertsOnASubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ascLocation = "centralus";
  const alertSimulatorRequestBody = {
    properties: {
      bundles: [
        "AppServices",
        "DNS",
        "KeyVaults",
        "KubernetesService",
        "ResourceManager",
        "SqlServers",
        "StorageAccounts",
        "VirtualMachines",
      ],
      kind: "Bundles",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.alerts.simulate(ascLocation, alertSimulatorRequestBody);
  console.log(result);
}

simulateSecurityAlertsOnASubscription().catch(console.error);
```
