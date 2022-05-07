Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-cognitiveservices_7.1.0/sdk/cognitiveservices/arm-cognitiveservices/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { CognitiveServicesManagementClient } = require("@azure/arm-cognitiveservices");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the state of specified commitmentPlans associated with the Cognitive Services account.
 *
 * @summary Update the state of specified commitmentPlans associated with the Cognitive Services account.
 * x-ms-original-file: specification/cognitiveservices/resource-manager/Microsoft.CognitiveServices/stable/2022-03-01/examples/PutCommitmentPlan.json
 */
async function putCommitmentPlan() {
  const subscriptionId = "subscriptionId";
  const resourceGroupName = "resourceGroupName";
  const accountName = "accountName";
  const commitmentPlanName = "commitmentPlanName";
  const commitmentPlan = {
    properties: {
      autoRenew: true,
      current: { tier: "T1" },
      hostingModel: "Web",
      planType: "Speech2Text",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CognitiveServicesManagementClient(credential, subscriptionId);
  const result = await client.commitmentPlans.createOrUpdate(
    resourceGroupName,
    accountName,
    commitmentPlanName,
    commitmentPlan
  );
  console.log(result);
}

putCommitmentPlan().catch(console.error);
```
