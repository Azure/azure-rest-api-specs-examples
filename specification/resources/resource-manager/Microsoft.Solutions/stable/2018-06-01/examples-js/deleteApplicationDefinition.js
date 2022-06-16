const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the managed application definition.
 *
 * @summary Deletes the managed application definition.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/deleteApplicationDefinition.json
 */
async function deleteApplicationDefinition() {
  const subscriptionId = "subid";
  const resourceGroupName = "rg";
  const applicationDefinitionName = "myManagedApplicationDef";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applicationDefinitions.beginDeleteByIdAndWait(
    resourceGroupName,
    applicationDefinitionName
  );
  console.log(result);
}

deleteApplicationDefinition().catch(console.error);
