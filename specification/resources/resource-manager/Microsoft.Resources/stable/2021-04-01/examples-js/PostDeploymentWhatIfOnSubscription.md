Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-resources_5.0.1/sdk/resources/arm-resources/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Returns changes that will be made by the deployment if executed at the scope of the subscription.
 *
 * @summary Returns changes that will be made by the deployment if executed at the scope of the subscription.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-04-01/examples/PostDeploymentWhatIfOnSubscription.json
 */
async function predictTemplateChangesAtSubscriptionScope() {
  const subscriptionId = "00000000-0000-0000-0000-000000000001";
  const deploymentName = "my-deployment";
  const parameters = {
    location: "westus",
    properties: { mode: "Incremental", parameters: {}, templateLink: {} },
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.deployments.beginWhatIfAtSubscriptionScopeAndWait(
    deploymentName,
    parameters
  );
  console.log(result);
}

predictTemplateChangesAtSubscriptionScope().catch(console.error);
```
