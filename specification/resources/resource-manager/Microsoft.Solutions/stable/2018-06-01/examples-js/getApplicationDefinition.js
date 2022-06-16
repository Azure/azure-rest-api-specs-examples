const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets the managed application definition.
 *
 * @summary Gets the managed application definition.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/getApplicationDefinition.json
 */
async function getManagedApplicationDefinition() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const applicationDefinitionName = "myManagedApplicationDef";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applicationDefinitions.get(
    resourceGroupName,
    applicationDefinitionName
  );
  console.log(result);
}

getManagedApplicationDefinition().catch(console.error);
