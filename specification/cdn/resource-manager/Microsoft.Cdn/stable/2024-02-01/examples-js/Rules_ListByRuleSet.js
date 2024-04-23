const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists all of the existing delivery rules within a rule set.
 *
 * @summary Lists all of the existing delivery rules within a rule set.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/Rules_ListByRuleSet.json
 */
async function rulesListByRuleSet() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const ruleSetName = "ruleSet1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.rules.listByRuleSet(resourceGroupName, profileName, ruleSetName)) {
    resArray.push(item);
  }
  console.log(resArray);
}
