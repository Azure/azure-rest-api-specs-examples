const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Lists environment types for the devcenter.
 *
 * @summary Lists environment types for the devcenter.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/EnvironmentTypes_List.json
 */
async function environmentTypesListByDevCenter() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const resArray = new Array();
  for await (let item of client.environmentTypes.listByDevCenter(
    resourceGroupName,
    devCenterName
  )) {
    resArray.push(item);
  }
  console.log(resArray);
}

environmentTypesListByDevCenter().catch(console.error);
