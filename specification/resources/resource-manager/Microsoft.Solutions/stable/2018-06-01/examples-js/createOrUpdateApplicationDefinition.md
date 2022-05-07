Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-managedapplications_2.0.1/sdk/managedapplications/arm-managedapplications/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new managed application definition.
 *
 * @summary Creates a new managed application definition.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/createOrUpdateApplicationDefinition.json
 */
async function createOrUpdateManagedApplicationDefinition() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const applicationDefinitionName = "myManagedApplicationDef";
  const parameters = {
    description: "myManagedApplicationDef description",
    authorizations: [{ principalId: "validprincipalguid", roleDefinitionId: "validroleguid" }],
    displayName: "myManagedApplicationDef",
    location: "East US 2",
    lockLevel: "None",
    packageFileUri: "https://path/to/packagezipfile",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applicationDefinitions.beginCreateOrUpdateByIdAndWait(
    resourceGroupName,
    applicationDefinitionName,
    parameters
  );
  console.log(result);
}

createOrUpdateManagedApplicationDefinition().catch(console.error);
```
