const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Partially updates an environment type.
 *
 * @summary Partially updates an environment type.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-08-01-preview/examples/EnvironmentTypes_Patch.json
 */
async function environmentTypesUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const environmentTypeName = "{environmentTypeName}";
  const body = { tags: { owner: "superuser" } };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.environmentTypes.update(
    resourceGroupName,
    devCenterName,
    environmentTypeName,
    body
  );
  console.log(result);
}

environmentTypesUpdate().catch(console.error);
