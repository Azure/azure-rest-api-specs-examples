Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the alert's state
 *
 * @summary Update the alert's state
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-11-01/examples/Alerts/UpdateAlertSubscriptionLocation_resolve_example.json
 */
async function updateSecurityAlertStateOnASubscriptionFromASecurityDataLocation() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ascLocation = "westeurope";
  const alertName = "2518298467986649999_4d25bfef-2d77-4a08-adc0-3e35715cc92a";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.alerts.updateSubscriptionLevelStateToResolve(ascLocation, alertName);
  console.log(result);
}

updateSecurityAlertStateOnASubscriptionFromASecurityDataLocation().catch(console.error);
```
