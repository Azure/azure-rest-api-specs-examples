const { ApiManagementClient } = require("@azure/arm-apimanagement");
const { DefaultAzureCredential } = require("@azure/identity");
require("dotenv/config");

/**
 * This sample demonstrates how to Creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 *
 * @summary Creates or updates an API Management service. This is long running operation and could take several minutes to complete.
 * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/ApiManagementCreateServiceWithCustomHostnameKeyVault.json
 */
async function apiManagementCreateServiceWithCustomHostnameKeyVault() {
  const subscriptionId =
    process.env["APIMANAGEMENT_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["APIMANAGEMENT_RESOURCE_GROUP"] || "rg1";
  const serviceName = "apimService1";
  const parameters = {
    apiVersionConstraint: { minApiVersion: "2019-01-01" },
    hostnameConfigurations: [
      {
        type: "Proxy",
        defaultSslBinding: true,
        hostName: "gateway1.msitesting.net",
        identityClientId: "329419bc-adec-4dce-9568-25a6d486e468",
        keyVaultId: "https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert",
      },
      {
        type: "Management",
        hostName: "mgmt.msitesting.net",
        identityClientId: "329419bc-adec-4dce-9568-25a6d486e468",
        keyVaultId: "https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert",
      },
      {
        type: "Portal",
        hostName: "portal1.msitesting.net",
        identityClientId: "329419bc-adec-4dce-9568-25a6d486e468",
        keyVaultId: "https://rpbvtkeyvaultintegration.vault.azure.net/secrets/msitestingCert",
      },
      {
        type: "ConfigurationApi",
        certificatePassword: "Password",
        encodedCertificate: "****** Base 64 Encoded Certificate ************",
        hostName: "configuration-api.msitesting.net",
      },
    ],
    identity: {
      type: "UserAssigned",
      userAssignedIdentities: {
        "/subscriptions/00000000000000000000000000000000/resourceGroups/rg1/providers/MicrosoftManagedIdentity/userAssignedIdentities/id1":
          {},
      },
    },
    location: "North Europe",
    publisherEmail: "apim@autorestsdk.com",
    publisherName: "autorestsdk",
    sku: { name: "Premium", capacity: 1 },
    tags: { tag1: "value1", tag2: "value2", tag3: "value3" },
    virtualNetworkType: "None",
  };
  const credential = new DefaultAzureCredential();
  const client = new ApiManagementClient(credential, subscriptionId);
  const result = await client.apiManagementService.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    parameters,
  );
  console.log(result);
}
