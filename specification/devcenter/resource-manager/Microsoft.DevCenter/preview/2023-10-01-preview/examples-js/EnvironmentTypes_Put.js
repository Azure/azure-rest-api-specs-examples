const { DevCenterClient } = require("@azure/arm-devcenter");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates an environment type.
 *
 * @summary Creates or updates an environment type.
 * x-ms-original-file: specification/devcenter/resource-manager/Microsoft.DevCenter/preview/2023-10-01-preview/examples/EnvironmentTypes_Put.json
 */
async function environmentTypesCreateOrUpdate() {
  const subscriptionId =
    process.env["DEVCENTER_SUBSCRIPTION_ID"] || "0ac520ee-14c0-480f-b6c9-0a90c58ffff";
  const resourceGroupName = process.env["DEVCENTER_RESOURCE_GROUP"] || "rg1";
  const devCenterName = "Contoso";
  const environmentTypeName = "DevTest";
  const body = {
    displayName: "Dev",
    tags: { owner: "superuser" },
  };
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
