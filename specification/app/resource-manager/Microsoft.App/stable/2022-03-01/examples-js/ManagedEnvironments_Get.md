Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appcontainers_1.0.0/sdk/appcontainers/arm-appcontainers/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { ContainerAppsAPIClient } = require("@azure/arm-appcontainers");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get the properties of a Managed Environment used to host container apps.
 *
 * @summary Get the properties of a Managed Environment used to host container apps.
 * x-ms-original-file: specification/app/resource-manager/Microsoft.App/stable/2022-03-01/examples/ManagedEnvironments_Get.json
 */
async function getEnvironmentsByName() {
  const subscriptionId = "8efdecc5-919e-44eb-b179-915dca89ebf9";
  const resourceGroupName = "examplerg";
  const environmentName = "jlaw-demo1";
  const credential = new DefaultAzureCredential();
  const client = new ContainerAppsAPIClient(credential, subscriptionId);
  const result = await client.managedEnvironments.get(resourceGroupName, environmentName);
  console.log(result);
}

getEnvironmentsByName().catch(console.error);
```
