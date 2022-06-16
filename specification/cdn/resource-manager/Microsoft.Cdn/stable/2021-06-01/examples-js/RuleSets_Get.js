const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets an existing AzureFrontDoor rule set with the specified rule set name under the specified subscription, resource group and profile.
 *
 * @summary Gets an existing AzureFrontDoor rule set with the specified rule set name under the specified subscription, resource group and profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2021-06-01/examples/RuleSets_Get.json
 */
async function ruleSetsGet() {
  const subscriptionId = "subid";
  const resourceGroupName = "RG";
  const profileName = "profile1";
  const ruleSetName = "ruleSet1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.ruleSets.get(resourceGroupName, profileName, ruleSetName);
  console.log(result);
}

ruleSetsGet().catch(console.error);
