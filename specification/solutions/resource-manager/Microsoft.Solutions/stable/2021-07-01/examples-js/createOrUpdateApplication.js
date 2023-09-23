const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates a managed application.
 *
 * @summary Creates or updates a managed application.
 * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/createOrUpdateApplication.json
 */
async function createOrUpdateManagedApplication() {
  const subscriptionId = process.env["MANAGEDAPPLICATIONS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["MANAGEDAPPLICATIONS_RESOURCE_GROUP"] || "rg";
  const applicationName = "myManagedApplication";
  const parameters = {
    applicationDefinitionId:
      "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef",
    kind: "ServiceCatalog",
    managedResourceGroupId: "/subscriptions/subid/resourceGroups/myManagedRG",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applications.beginCreateOrUpdateAndWait(
    resourceGroupName,
    applicationName,
    parameters
  );
  console.log(result);
}
