const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes a Dev Box definition
 *
 * @summary Deletes a Dev Box definition
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/DevBoxDefinitions_Delete.json
 */
async function devBoxDefinitionsDelete() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const devCenterName = "Contoso";
  const devBoxDefinitionName = "WebDevBox";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.devBoxDefinitions.beginDeleteAndWait(
    resourceGroupName,
    devCenterName,
    devBoxDefinitionName
  );
  console.log(result);
}
