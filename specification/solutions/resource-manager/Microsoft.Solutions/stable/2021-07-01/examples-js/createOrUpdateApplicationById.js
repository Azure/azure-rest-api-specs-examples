const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a managed application.
 *
 * @summary Creates or updates a managed application.
 * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateApplicationById.json
 */
async function createsOrUpdatesAManagedApplication() {
  const applicationId =
    "subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applications/myManagedApplication";
  const parameters = {
    applicationDefinitionId:
      "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef",
    kind: "ServiceCatalog",
    managedResourceGroupId: "/subscriptions/subid/resourceGroups/myManagedRG",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential);
  const result = await client.applications.beginCreateOrUpdateByIdAndWait(
    applicationId,
    parameters
  );
  console.log(result);
}
