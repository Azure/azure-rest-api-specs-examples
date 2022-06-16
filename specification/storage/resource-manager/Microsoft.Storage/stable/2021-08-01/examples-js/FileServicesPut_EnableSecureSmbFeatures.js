const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

async function putFileServicesEnableSecureSmbFeatures() {
  const subscriptionId = "{subscription-id}";
  const resourceGroupName = "res4410";
  const accountName = "sto8607";
  const parameters = {
    protocolSettings: {
      smb: {
        authenticationMethods: "NTLMv2;Kerberos",
        channelEncryption: "AES-128-CCM;AES-128-GCM;AES-256-GCM",
        kerberosTicketEncryption: "RC4-HMAC;AES-256",
        versions: "SMB2.1;SMB3.0;SMB3.1.1",
      },
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new StorageManagementClient(credential, subscriptionId);
  const result = await client.fileServices.setServiceProperties(
    resourceGroupName,
    accountName,
    parameters
  );
  console.log(result);
}

putFileServicesEnableSecureSmbFeatures().catch(console.error);
