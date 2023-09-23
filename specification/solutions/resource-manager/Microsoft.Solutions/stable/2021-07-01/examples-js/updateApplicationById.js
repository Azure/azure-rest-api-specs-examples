const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing managed application.
 *
 * @summary Updates an existing managed application.
 * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/updateApplicationById.json
 */
async function updatesAnExistingManagedApplication() {
  const applicationId =
    "subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication";
  const parameters = {
    applicationDefinitionId:
      "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef",
    kind: "ServiceCatalog",
    managedResourceGroupId: "/subscriptions/subid/resourceGroups/myManagedRG",
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential);
  const result = await client.applications.beginUpdateByIdAndWait(applicationId, options);
  console.log(result);
}
