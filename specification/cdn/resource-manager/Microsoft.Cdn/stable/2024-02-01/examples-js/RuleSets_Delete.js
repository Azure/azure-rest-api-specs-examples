const { CdnManagementClient } = require("@azure/arm-cdn");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an existing AzureFrontDoor rule set with the specified rule set name under the specified subscription, resource group and profile.
 *
 * @summary Deletes an existing AzureFrontDoor rule set with the specified rule set name under the specified subscription, resource group and profile.
 * x-ms-original-file: specification/cdn/resource-manager/Microsoft.Cdn/stable/2024-02-01/examples/RuleSets_Delete.json
 */
async function ruleSetsDelete() {
  const subscriptionId = process.env["CDN_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["CDN_RESOURCE_GROUP"] || "RG";
  const profileName = "profile1";
  const ruleSetName = "ruleSet1";
  const credential = new DefaultAzureCredential();
  const client = new CdnManagementClient(credential, subscriptionId);
  const result = await client.ruleSets.beginDeleteAndWait(
    resourceGroupName,
    profileName,
    ruleSetName,
  );
  console.log(result);
}
