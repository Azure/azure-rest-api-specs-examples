Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-devtestlabs_4.0.1/sdk/devtestlabs/arm-devtestlabs/README.md) on how to add the SDK to your project and authenticate.

```javascript
const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or replace an existing policy.
 *
 * @summary Create or replace an existing policy.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Policies_CreateOrUpdate.json
 */
async function policiesCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const policySetName = "{policySetName}";
  const name = "{policyName}";
  const policy = {
    description: "{policyDescription}",
    evaluatorType: "{policyEvaluatorType}",
    factData: "{policyFactData}",
    factName: "{policyFactName}",
    location: "{location}",
    status: "{policyStatus}",
    tags: { tagName1: "tagValue1" },
    threshold: "{policyThreshold}",
  };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.policies.createOrUpdate(
    resourceGroupName,
    labName,
    policySetName,
    name,
    policy
  );
  console.log(result);
}

policiesCreateOrUpdate().catch(console.error);
```
