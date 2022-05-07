Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Get all Kubernetes Environments for a subscription.
 *
 * @summary Description for Get all Kubernetes Environments for a subscription.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/KubeEnvironments_ListBySubscription.json
 */
async function listKubeEnvironmentsBySubscription() {
  const subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.kubeEnvironments.listBySubscription()) {
    resArray.push(item);
  }
  console.log(resArray);
}

listKubeEnvironmentsBySubscription().catch(console.error);
```
