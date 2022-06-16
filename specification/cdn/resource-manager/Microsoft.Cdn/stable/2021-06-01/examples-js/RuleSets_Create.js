const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates a new rule set within the specified profile.
 *
 * @summary Creates a new rule set within the specified profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/RuleSets_Create.json
 */
async function ruleSetsCreate() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const ruleSetName = "ruleSet1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.ruleSets.create(resourceGroupName, profileName, ruleSetName);
  console.log(result);
}

ruleSetsCreate().catch(console.error);
