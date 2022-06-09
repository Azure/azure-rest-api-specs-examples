```javascript
const { StorageManagementClient } = require("@azure/arm-storage");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Sets the properties of file services in storage accounts, including CORS (Cross-Origin Resource Sharing) rules.
 *
 * @summary Sets the properties of file services in storage accounts, including CORS (Cross-Origin Resource Sharing) rules.
 * x-ms-original-file: specification/storage/resource-manager/Microsoft.Storage/stable/2021-09-01/examples/FileServicesPut_EnableSecureSmbFeatures.json
 */
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-js/blob/%40azure%2Farm-storage_17.2.0/sdk/storage/arm-storage/README.md) on how to add the SDK to your project and authenticate.
