```javascript
const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the datadog single sign-on resource for the given Monitor.
 *
 * @summary Gets the datadog single sign-on resource for the given Monitor.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/SingleSignOnConfigurations_Get.json
 */
async function singleSignOnConfigurationsGet() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const configurationName = "default";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const result = await client.singleSignOnConfigurations.get(
    resourceGroupName,
    monitorName,
    configurationName
  );
  console.log(result);
}

singleSignOnConfigurationsGet().catch(console.error);
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-datadog_3.0.1/sdk/datadog/arm-datadog/README.md) on how to add the SDK to your project and authenticate.
