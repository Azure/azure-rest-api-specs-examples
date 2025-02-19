const { CosmosDBManagementClient } = require("@azure/arm-cosmosdb");
const { DefaultAzureCredential } = require("@azure/identity");

/**
 * This sample demonstrates how to Create or update a ClientEncryptionKey. This API is meant to be invoked via tools such as the Azure Powershell (instead of directly).
 *
 * @summary Create or update a ClientEncryptionKey. This API is meant to be invoked via tools such as the Azure Powershell (instead of directly).
 * x-ms-original-file: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBSqlClientEncryptionKeyCreateUpdate.json
 */
async function cosmosDbClientEncryptionKeyCreateUpdate() {
  const subscriptionId = process.env["COSMOSDB_SUBSCRIPTION_ID"] || "subId";
  const resourceGroupName = process.env["COSMOSDB_RESOURCE_GROUP"] || "rgName";
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
        "VGhpcyBpcyBhY3R1YWxseSBhbiBhcnJheSBvZiBieXRlcy4gVGhpcyByZXF1ZXN0L3Jlc3BvbnNlIGlzIGJlaW5nIHByZXNlbnRlZCBhcyBhIHN0cmluZyBmb3IgcmVhZGFiaWxpdHkgaW4gdGhlIGV4YW1wbGU=",
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
    createUpdateClientEncryptionKeyParameters,
  );
  console.log(result);
}
