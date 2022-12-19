const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Deletes an environment type.
 *
 * @summary Deletes an environment type.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-11-11-preview/examples/EnvironmentTypes_Delete.json
 */
async function environmentTypesDelete() {
  const subscriptionId = "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const environmentTypeName = "DevTest";
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.environmentTypes.delete(
    resourceGroupName,
    devCenterName,
    environmentTypeName
  );
  console.log(result);
}

environmentTypesDelete().catch(console.error);
