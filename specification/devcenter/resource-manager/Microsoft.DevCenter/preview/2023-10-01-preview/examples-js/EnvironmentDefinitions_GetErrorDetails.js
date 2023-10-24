const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets Environment Definition error details
 *
 * @summary Gets Environment Definition error details
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/EnvironmentDefinitions_GetErrorDetails.json
 */
async function environmentDefinitionsGetErrorDetails() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const devCenterName = "Contoso";
  const catalogName = "myCatalog";
  const environmentDefinitionName = "myEnvironmentDefinition";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.environmentDefinitions.getErrorDetails(
    resourceGroupName,
    devCenterName,
    catalogName,
    environmentDefinitionName
  );
  console.log(result);
}
