const { AppPlatformManagementClient } = require("@azure/arm-appplatform");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update certificate resource.
 *
 * @summary Create or update certificate resource.
 * x-ms-original-file: specification/appplatform/resource-manager/Microsoft.AppPlatform/stable/2022-04-01/examples/Certificates_CreateOrUpdate.json
 */
async function certificatesCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const serviceName = "myservice";
  const certificateName = "mycertificate";
  const certificateResource = {
    properties: {
      type: "KeyVaultCertificate",
      certVersion: "08a219d06d874795a96db47e06fbb01e",
      keyVaultCertName: "mycert",
      vaultUri: "https://myvault.vault.azure.net",
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new AppPlatformManagementClient(credential, subscriptionId);
  const result = await client.certificates.beginCreateOrUpdateAndWait(
    resourceGroupName,
    serviceName,
    certificateName,
    certificateResource
  );
  console.log(result);
}

certificatesCreateOrUpdate().catch(console.error);
