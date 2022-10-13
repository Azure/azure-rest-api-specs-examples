const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a ClientEncryptionKey. This API is meant to be invoked via tools such as the Azure Powershell (instead of directly).
 *
 * @summary Create or update a ClientEncryptionKey. This API is meant to be invoked via tools such as the Azure Powershell (instead of directly).
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2022-08-15-preview/examples/CosmosDBSqlClientEncryptionKeyCreateUpdate.json
 */
async function cosmosDbClientEncryptionKeyCreateUpdate() {
  const subscriptionId = "subId";
  const resourceGroupName = "rgName";
  const accountName = "accountName";
  const databaseName = "databaseName";
  const clientEncryptionKeyName = "cekName";
  const createUpdateClientEncryptionKeyParameters = {
    resource: {
      encryptionAlgorithm: "AEAD_AES_256_CBC_HMAC_SHA256",
      id: "cekName",
      keyWrapMetadata: {
        name: "customerManagedKey",
        type: "AzureKeyVault",
        algorithm: "RSA-OAEP",
        value: "AzureKeyVault Key URL",
      },
      wrappedDataEncryptionKey: Buffer.from(
        "This is actually an array of bytes. This request/response is being presented as a string for readability in the example"
      ),
    },
  };
  const credential = new DefaultAzureCredential();
  const client = new CosmosDBManagementClient(credential, subscriptionId);
  const result = await client.sqlResources.beginCreateUpdateClientEncryptionKeyAndWait(
    resourceGroupName,
    accountName,
    databaseName,
    clientEncryptionKeyName,
    createUpdateClientEncryptionKeyParameters
  );
  console.log(result);
}

cosmosDbClientEncryptionKeyCreateUpdate().catch(console.error);
