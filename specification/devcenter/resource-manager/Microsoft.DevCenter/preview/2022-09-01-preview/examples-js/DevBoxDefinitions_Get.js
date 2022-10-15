const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Gets a Dev Box definition
 *
 * @summary Gets a Dev Box definition
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/DevBoxDefinitions_Get.json
 */
async function devBoxDefinitionsGet() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const devBoxDefinitionName = "WebDevBox";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.devBoxDefinitions.get(
    resourceGroupName,
    devCenterName,
    devBoxDefinitionName
  );
  console.log(result);
}

devBoxDefinitionsGet().catch(console.error);
