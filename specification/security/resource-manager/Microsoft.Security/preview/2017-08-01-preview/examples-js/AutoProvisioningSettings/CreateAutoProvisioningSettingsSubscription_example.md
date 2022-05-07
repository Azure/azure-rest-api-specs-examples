Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Details of a specific setting
 *
 * @summary Details of a specific setting
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/AutoProvisioningSettings/CreateAutoProvisioningSettingsSubscription_example.json
 */
async function createAutoProvisioningSettingsForSubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const settingName = "default";
  const setting = {
    name: "default",
    type: "Microsoft.Security/autoProvisioningSettings",
    autoProvision: "On",
    id: "/subscriptions/20ff7fc3-e762-44dd-bd96-b71116dcdc23/providers/Microsoft.Security/autoProvisioningSettings/default",
  };
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.autoProvisioningSettings.create(settingName, setting);
  console.log(result);
}

createAutoProvisioningSettingsForSubscription().catch(console.error);
```
