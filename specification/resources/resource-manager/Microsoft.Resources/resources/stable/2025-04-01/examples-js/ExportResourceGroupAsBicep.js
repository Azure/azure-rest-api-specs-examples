const { ResourceManagementClient } = require("@azure/arm-resources");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to captures the specified resource group as a template.
 *
 * @summary captures the specified resource group as a template.
 * x-ms-original-file: 2025-04-01/ExportResourceGroupAsBicep.json
 */
async function exportAResourceGroupAsBicep() {
  const credential = new DefaultAzureCredential();
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const client = new ResourceManagementClient(credential, subscriptionId);
  const result = await client.resourceGroups.exportTemplate("my-resource-group", {
    options: "IncludeParameterDefaultValue,IncludeComments",
    outputFormat: "Bicep",
    resources: ["*"],
  });
  console.log(result);
}
