const { ApplicationClient } = require("@azure/arm-managedapplications");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Updates an existing managed application.
 *
 * @summary Updates an existing managed application.
 * x-ms-original-file: specification/solutions/resource-manager/Microsoft.Solutions/stable/2021-07-01/examples/updateApplication.json
 */
async function updatesManagedApplication() {
  const subscriptionId = process.env["MANAGEDAPPLICATIONS_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["MANAGEDAPPLICATIONS_RESOURCE_GROUP"] || "rg";
  const applicationName = "myManagedApplication";
  const parameters = {
    applicationDefinitionId:
      "/subscriptions/subid/resourceGroups/rg/providers/Microsoft.Solutions/applicationDefinitions/myAppDef",
    kind: "ServiceCatalog",
    managedResourceGroupId: "/subscriptions/subid/resourceGroups/myManagedRG",
  };
  const options = { parameters };
  const credential = new DefaultAzureCredential();
  const client = new ApplicationClient(credential, subscriptionId);
  const result = await client.applications.beginUpdateAndWait(
    resourceGroupName,
    applicationName,
    options
  );
  console.log(result);
}
