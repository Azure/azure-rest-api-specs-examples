Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create setting for ingesting security data and logs to correlate with resources associated with the subscription.
 *
 * @summary Create setting for ingesting security data and logs to correlate with resources associated with the subscription.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-01-15-preview/examples/IngestionSettings/CreateIngestionSetting_example.json
 */
async function createAnIngestionSettingForSubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const ingestionSettingName = "default";
  const ingestionSetting = {};
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.ingestionSettings.create(ingestionSettingName, ingestionSetting);
  console.log(result);
}

createAnIngestionSettingForSubscription().catch(console.error);
```
