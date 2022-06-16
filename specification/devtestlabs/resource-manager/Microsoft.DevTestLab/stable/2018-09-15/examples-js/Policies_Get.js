const { DevTestLabsClient } = require("@azure/arm-devtestlabs");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Get policy.
 *
 * @summary Get policy.
 * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Policies_Get.json
 */
async function policiesGet() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "resourceGroupName";
  const labName = "{labName}";
  const policySetName = "{policySetName}";
  const name = "{policyName}";
  const credential = new DefaultAzureCredential();
  const client = new DevTestLabsClient(credential, subscriptionId);
  const result = await client.policies.get(resourceGroupName, labName, policySetName, name);
  console.log(result);
}

policiesGet().catch(console.error);
