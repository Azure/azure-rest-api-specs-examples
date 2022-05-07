Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-security_5.0.0/sdk/security/arm-security/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { SecurityCenter } = require("@azure/arm-security");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a single custom entity store assignment by name for the provided subscription and resource group.
 *
 * @summary Gets a single custom entity store assignment by name for the provided subscription and resource group.
 * x-ms-original-file: specification/security/resource-manager/Microsoft.Security/preview/2021-07-01-preview/examples/CustomEntityStoreAssignments/customEntityStoreAssignmentGet_example.json
 */
async function getACustomEntityStoreAssignment() {
  const subscriptionId = "e5d1b86c-3051-44d5-8802-aa65d45a279b";
  const resourceGroupName = "TestResourceGroup";
  const customEntityStoreAssignmentName = "33e7cc6e-a139-4723-a0e5-76993aee0771";
  const credential = new DefaultAzureCredential();
  const client = new SecurityCenter(credential, subscriptionId);
  const result = await client.customEntityStoreAssignments.get(
    resourceGroupName,
    customEntityStoreAssignmentName
  );
  console.log(result);
}

getACustomEntityStoreAssignment().catch(console.error);
```
