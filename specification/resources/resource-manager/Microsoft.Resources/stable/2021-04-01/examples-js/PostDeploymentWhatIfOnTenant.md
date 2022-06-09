```javascript
const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns changes that will be made by the deployment if executed at the scope of the tenant group.
 *
 * @summary Returns changes that will be made by the deployment if executed at the scope of the tenant group.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnTenant.json
 */
async function predictTemplateChangesAtManagementGroupScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const deploymentName = "exampleDeploymentName";
  const parameters = {
    location: "eastus",
    properties: { mode: "Incremental", parameters: {}, templateLink: {} },
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginWhatIfAtTenantScopeAndWait(
    deploymentName,
    parameters
  );
  console.log(result);
}

predictTemplateChangesAtManagementGroupScope().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-resources_5.0.1/sdk/resources/arm-resources/README.md) on how to add the SDK to your project and authenticate.
