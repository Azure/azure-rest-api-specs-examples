const { NginxManagementClient } = require("@azure/arm-nginx");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update the NGINX certificates for given NGINX deployment
 *
 * @summary Create or update the NGINX certificates for given NGINX deployment
 * x-ms-original-file: specification/nginx/resource-manager/NGINX.NGINXPLUS/preview/2024-01-01-preview/examples/Certificates_CreateOrUpdate.json
 */
async function certificatesCreateOrUpdate() {
  const subscriptionId =
    process.env["NGINX_SUBSCRIPTION_ID"] || "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = process.env["NGINX_RESOURCE_GROUP"] || "myResourceGroup";
  const deploymentName = "myDeployment";
  const certificateName = "default";
  const body = {
    properties: {
      certificateVirtualPath: "/src/cert/somePath.cert",
      keyVaultSecretId: "https://someKV.vault.azure.com/someSecretID",
      keyVirtualPath: "/src/cert/somekey.key",
    },
  };
  const options = { body };
  const credential = new DefaultAzureCredential();
  const client = new NginxManagementClient(credential, subscriptionId);
  const result = await client.certificates.beginCreateOrUpdateAndWait(
    resourceGroupName,
    deploymentName,
    certificateName,
    options,
  );
  console.log(result);
}
