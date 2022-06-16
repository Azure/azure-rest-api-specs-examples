const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing managed application. The only value that can be updated via PATCH currently is the tags.
 *
 * @summary Updates an existing managed application. The only value that can be updated via PATCH currently is the tags.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Solutions/stable/2018-06-01/examples/updateApplicationById.json
 */
async function updateApplicationById() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const applicationId = "myApplicationId";
  const parameters = {
    applicationDefinitionId:
      "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef",
    kind: "ServiceCatalog",
    managedResourceGroupId: "/subscriptions/subid/resourceGroups/myManagedRG",
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applications.updateById(applicationId, options);
  console.log(result);
}

updateApplicationById().catch(console.error);
