Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managedapplications_2.0.1/sdk/managedapplications/arm-managedapplications/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the managed application definition.
 *
 * @summary Gets the managed application definition.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/getApplicationDefinition.json
 */
async function getManagedApplicationDefinition() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const applicationDefinitionName = "myManagedApplicationDef";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applicationDefinitions.get(
    resourceGroupName,
    applicationDefinitionName
  );
  console.log(result);
}

getManagedApplicationDefinition().catch(console.error);
```
