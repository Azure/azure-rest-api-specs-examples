Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Settings about where we should store your security data and logs. If the result is empty, it means that no custom-workspace configuration was set
 *
 * @summary Settings about where we should store your security data and logs. If the result is empty, it means that no custom-workspace configuration was set
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2017-08-01-preview/examples/WorkspaceSettings/GetWorkspaceSetting_example.json
 */
async function getAWorkspaceSettingOnSubscription() {
  const subscriptionId = "20ff7fc3-e762-44dd-bd96-b71116dcdc23";
  const workspaceSettingName = "default";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.workspaceSettings.get(workspaceSettingName);
  console.log(result);
}

getAWorkspaceSettingOnSubscription().catch(console.error);
```
