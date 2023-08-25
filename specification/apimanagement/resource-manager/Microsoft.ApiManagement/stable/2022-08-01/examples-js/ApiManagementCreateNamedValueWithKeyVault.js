const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Creates or updates named value.
 *
 * @summary Creates or updates named value.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateNamedValueWithKeyVault.json
 */
async function apiManagementCreateNamedValueWithKeyVault() {
  const subscriptionId = process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "subid";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const namedValueId = "testprop6";
  const parameters = {
    displayName: "prop6namekv",
    keyVault: {
      identityClientId: "ceaa6b06-c00f-43ef-99ac-f53d1fe876a0",
      secretIdentifier: "https://contoso.vault.azure.net/secrets/aadSecret",
    },
    secret: true,
    tags: ["foo", "bar"],
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.namedValue.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    namedValueId,
    parameters
  );
  console.log(result);
}
