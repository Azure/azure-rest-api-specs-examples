const { ExternalIdentitiesConfigurationClient } = require("@azure/arm-azureadexternalidentities");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Update the Azure AD B2C tenant resource.
 *
 * @summary Update the Azure AD B2C tenant resource.
 * x-ms-original-file: specification/cpim/resource-manager/Microsoft.AzureActiveDirectory/stable/2021-04-01/examples/updateTenant.json
 */
async function updateTenant() {
  const subscriptionId = "34adfa4f-cedf-4dc0-ba29-b6d1a69ab345";
  const resourceGroupName = "contosoResourceGroup";
  const resourceName = "contoso.onmicrosoft.com";
  const updateTenantRequestBody = {
    billingConfig: { billingType: "MAU" },
    sku: { name: "PremiumP1" },
    tags: { key: "value" },
  };
  const options = { updateTenantRequestBody };
  const credential = new DefaultAzureCredential();
  const client = new ExternalIdentitiesConfigurationClient(credential, subscriptionId);
  const result = await client.b2CTenants.update(resourceGroupName, resourceName, options);
  console.log(result);
}

updateTenant().catch(console.error);
