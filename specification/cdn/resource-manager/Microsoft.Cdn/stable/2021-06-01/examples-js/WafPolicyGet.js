const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Retrieve protection policy with specified name within a resource group.
 *
 * @summary Retrieve protection policy with specified name within a resource group.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/WafPolicyGet.json
 */
async function getPolicy() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg1";
  const policyName = "MicrosoftCdnWafPolicy";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.policies.get(resourceGroupName, policyName);
  console.log(result);
}

getPolicy().catch(console.error);
