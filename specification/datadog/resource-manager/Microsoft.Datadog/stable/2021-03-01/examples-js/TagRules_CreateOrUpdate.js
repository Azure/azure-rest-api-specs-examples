const { MicrosoftDatadogClient } = require("@azure/arm-datadog");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a tag rule set for a given monitor resource.
 *
 * @summary Create or update a tag rule set for a given monitor resource.
 * x-ms-original-file: specification/datadog/resource-manager/Microsoft.Datadog/stable/2021-03-01/examples/TagRules_CreateOrUpdate.json
 */
async function tagRulesCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const monitorName = "myMonitor";
  const ruleSetName = "default";
  const credential = new DefaultAzureCredential();
  const client = new MicrosoftDatadogClient(credential, subscriptionId);
  const result = await client.tagRules.createOrUpdate(resourceGroupName, monitorName, ruleSetName);
  console.log(result);
}

tagRulesCreateOrUpdate().catch(console.error);
