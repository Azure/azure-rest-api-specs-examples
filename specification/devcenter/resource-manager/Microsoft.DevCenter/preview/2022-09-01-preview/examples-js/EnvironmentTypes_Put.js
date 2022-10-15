const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an environment type.
 *
 * @summary Creates or updates an environment type.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2022-09-01-preview/examples/EnvironmentTypes_Put.json
 */
async function environmentTypesCreateOrUpdate() {
  const subscriptionId = "{subscriptionId}";
  const resourceGroupName = "rg1";
  const devCenterName = "Contoso";
  const environmentTypeName = "{environmentTypeName}";
  const body = { tags: { owner: "superuser" } };
  const credential = new DefaultAzureCredential();
  const client = new DevCenterClient(credential, subscriptionId);
  const result = await client.environmentTypes.createOrUpdate(
    resourceGroupName,
    devCenterName,
    environmentTypeName,
    body
  );
  console.log(result);
}

environmentTypesCreateOrUpdate().catch(console.error);
