const { SignalRManagementClient } = require("@azure/arm-signalr");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a custom certificate.
 *
 * @summary Create or update a custom certificate.
 * x-ms-original-file: specification/signalr/resource-manager/Microsoft.SignalRService/stable/2022-02-01/examples/SignalRCustomCertificates_CreateOrUpdate.json
 */
async function signalRCustomCertificatesCreateOrUpdate() {
  const subscriptionId = "00000000-0000-0000-0000-000000000000";
  const resourceGroupName = "myResourceGroup";
  const resourceName = "mySignalRService";
  const certificateName = "myCert";
  const parameters = {
    keyVaultBaseUri: "https://myvault.keyvault.azure.net/",
    keyVaultSecretName: "mycert",
    keyVaultSecretVersion: "bb6a44b2743f47f68dad0d6cc9756432",
  };
  const credential = new DefaultAzureCredential();
  const client = new SignalRManagementClient(credential, subscriptionId);
  const result = await client.signalRCustomCertificates.beginCreateOrUpdateAndWait(
    resourceGroupName,
    resourceName,
    certificateName,
    parameters
  );
  console.log(result);
}

signalRCustomCertificatesCreateOrUpdate().catch(console.error);
