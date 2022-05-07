Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-operationalinsights_8.0.1/sdk/operationalinsights/arm-operationalinsights/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { OperationalInsightsManagementClient } = require("@azure/arm-operationalinsights");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Enables an intelligence pack for a given workspace.
 *
 * @summary Enables an intelligence pack for a given workspace.
 * x-ms-original-file: specification/operationalinsights/resource-manager/Microsoft.OperationalInsights/stable/2020-08-01/examples/WorkspacesEnableIntelligencePack.json
 */
async function intelligencePacksEnable() {
  const subscriptionId = "00000000-0000-0000-0000-00000000000";
  const resourceGroupName = "rg1";
  const workspaceName = "TestLinkWS";
  const intelligencePackName = "ChangeTracking";
  const credential = new DefaultAzureCredential();
  const client = new OperationalInsightsManagementClient(credential, subscriptionId);
  const result = await client.intelligencePacks.enable(
    resourceGroupName,
    workspaceName,
    intelligencePackName
  );
  console.log(result);
}

intelligencePacksEnable().catch(console.error);
```
