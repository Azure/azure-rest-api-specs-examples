const { ResourceManagementClient } = require("@azure/arm-resources-profile-2020-09-01-hybrid");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Captures the specified resource group as a template.
 *
 * @summary Captures the specified resource group as a template.
 * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2019-10-01/examples/ExportResourceGroupWithFiltering.json
 */
async function exportAResourceGroupWithFiltering() {
  const subscriptionId =
    process.env["RESOURCES_SUBSCRIPTION_ID"] || "eaee6a92-e973-4922-9471-3a0a6abf81cd";
  const resourceGroupName = process.env["RESOURCES_RESOURCE_GROUP"] || "myResourceGroup";
  const parameters = {
    options: "SkipResourceNameParameterization",
    resources: [
      "/subscriptions/eaee6a92-e973-4922-9471-3a0a6abf81cd/resourceGroups/myResourceGroup/providers/My.RP/myResourceType/myFirstResource",
    ],
  };
  const credential = new DefaultAzureCredential();
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.resourceGroups.beginExportTemplateAndWait(
    resourceGroupName,
    parameters
  );
  console.log(result);
}
