Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-appservice_12.0.0/sdk/appservice/arm-appservice/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { WebSiteManagementClient } = require("@azure/arm-appservice");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Description for Start capturing network packets for the site.
 *
 * @summary Description for Start capturing network packets for the site.
 * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/StartWebSiteNetworkTraceOperation.json
 */
async function startANewNetworkTraceOperationForASite() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "testrg123";
  const name = "SampleApp";
  const durationInSeconds = 60;
  const options = { durationInSeconds };
  const credential = new DefaultAzureCredential();
  const client = new WebSiteManagementClient(credential, subscriptionId);
  const result = await client.webApps.beginStartNetworkTraceAndWait(
    resourceGroupName,
    name,
    options
  );
  console.log(result);
}

startANewNetworkTraceOperationForASite().catch(console.error);
```
