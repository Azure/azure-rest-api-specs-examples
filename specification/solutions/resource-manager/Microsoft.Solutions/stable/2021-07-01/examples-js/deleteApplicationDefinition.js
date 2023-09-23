const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes the managed application definition.
 *
 * @summary Deletes the managed application definition.
 * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/deleteApplicationDefinition.json
 */
async function deletesManagedApplicationDefinition() {
  const subscriptionId = process.env["MANAGEDAPPLICATIONS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["MANAGEDAPPLICATIONS_RESOURCE_GROUP"] || "rg";
  const applicationDefinitionName = "myManagedApplicationDef";
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applicationDefinitions.deleteById(
    resourceGroupName,
    applicationDefinitionName
  );
  console.log(result);
}
