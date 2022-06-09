```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Settings of different configurations in security center
 *
 * @summary Settings of different configurations in security center
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/stable/2021-07-01/examples/Settings/GetSetting_example.json
 */
async function getASettingOnSubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const settingName = "MCAS";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.settings.get(settingName);
  console.log(result);
}

getASettingOnSubscription().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.
