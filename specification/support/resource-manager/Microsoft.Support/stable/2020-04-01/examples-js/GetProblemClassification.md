Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-support_2.0.1/sdk/support/arm-support/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { MicrosoftSupport } = require("@azure/arm-support");
const { DefaultAzureCredential } = require("@azure/identity");

async function getsDetailsOfProblemClassificationForAzureService() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const serviceName = "service_guid";
  const problemClassificationName = "problemClassification_guid";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftSupport(credential, subscriptionId);
  const result = await client.problemClassifications.get(serviceName, problemClassificationName);
  console.log(result);
}

getsDetailsOfProblemClassificationForAzureService().catch(console.error);
```
