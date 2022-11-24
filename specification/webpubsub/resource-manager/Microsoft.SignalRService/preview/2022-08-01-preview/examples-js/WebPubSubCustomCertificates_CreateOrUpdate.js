const { WebPubSubManagementClient } = require("@azure/arm-webpubsub");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a custom certificate.
 *
 * @summary Create or update a custom certificate.
 * x-ms-original-file: specification/webpubsub/resource-manager/Microsoft.SignalRService/preview/2022-08-01-preview/examples/WebPubSubCustomCertificates_CreateOrUpdate.json
 */
async function webPubSubCustomCertificatesCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "myWebPubSubService";
  const certificateName = "myCert";
  const parameters = {
    keyVaultBaseUri: "https://myvault.keyvault.azure.net/",
    keyVaultSecretName: "mycert",
    keyVaultSecretVersion: "bb6a44b2743f47f68dad0d6cc9756432",
  };
  const credential = new DefaultAzureCredential();
  const client = new WebPubSubManagementClient(credential, subscriptionId);
  const result = await client.webPubSubCustomCertificates.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    certificateName,
    parameters
  );
  console.log(result);
}

webPubSubCustomCertificatesCreateOrUpdate().catch(console.error);
