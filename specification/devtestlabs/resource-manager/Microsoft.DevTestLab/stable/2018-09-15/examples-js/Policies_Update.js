const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Allows modifying tags of policies. All other properties will be ignored.
 *
 * @summary Allows modifying tags of policies. All other properties will be ignored.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Policies_Update.json
 */
async function policiesUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const policySetName = "{policySetName}";
  const name = "{policyName}";
  const policy = { tags: { tagName1: "tagValue1" } };
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.policies.update(
    resourceGroupName,
    labName,
    policySetName,
    name,
    policy
  );
  console.log(result);
}

policiesUpdate().catch(console.error);
