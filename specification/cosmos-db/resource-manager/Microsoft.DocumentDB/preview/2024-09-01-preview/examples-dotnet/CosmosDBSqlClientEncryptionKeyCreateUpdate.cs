using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.CosmosDB.Models;
using Azure.ResourceManager.CosmosDB;

// Generated from example definition: specification/cosmos-db/resource-manager/Microsoft.DocumentDB/preview/2024-09-01-preview/examples/CosmosDBSqlClientEncryptionKeyCreateUpdate.json
// this example is just showing the usage of "SqlResources_CreateUpdateClientEncryptionKey" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CosmosDBSqlDatabaseResource created on azure
// for more information of creating CosmosDBSqlDatabaseResource, please refer to the document of CosmosDBSqlDatabaseResource
string subscriptionId = "subId";
string resourceGroupName = "rgName";
string accountName = "accountName";
string databaseName = "databaseName";
ResourceIdentifier cosmosDBSqlDatabaseResourceId = CosmosDBSqlDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, databaseName);
CosmosDBSqlDatabaseResource cosmosDBSqlDatabase = client.GetCosmosDBSqlDatabaseResource(cosmosDBSqlDatabaseResourceId);

// get the collection of this CosmosDBSqlClientEncryptionKeyResource
CosmosDBSqlClientEncryptionKeyCollection collection = cosmosDBSqlDatabase.GetCosmosDBSqlClientEncryptionKeys();

// invoke the operation
string clientEncryptionKeyName = "cekName";
CosmosDBSqlClientEncryptionKeyCreateOrUpdateContent content = new CosmosDBSqlClientEncryptionKeyCreateOrUpdateContent(new CosmosDBSqlClientEncryptionKeyResourceInfo()
{
    Id = "cekName",
    EncryptionAlgorithm = "AEAD_AES_256_CBC_HMAC_SHA256",
    WrappedDataEncryptionKey = Convert.FromBase64String("VGhpcyBpcyBhY3R1YWxseSBhbiBhcnJheSBvZiBieXRlcy4gVGhpcyByZXF1ZXN0L3Jlc3BvbnNlIGlzIGJlaW5nIHByZXNlbnRlZCBhcyBhIHN0cmluZyBmb3IgcmVhZGFiaWxpdHkgaW4gdGhlIGV4YW1wbGU="),
    KeyWrapMetadata = new CosmosDBKeyWrapMetadata()
    {
        Name = "customerManagedKey",
        CosmosDBKeyWrapMetadataType = "AzureKeyVault",
        Value = "AzureKeyVault Key URL",
        Algorithm = "RSA-OAEP",
    },
});
ArmOperation<CosmosDBSqlClientEncryptionKeyResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, clientEncryptionKeyName, content);
CosmosDBSqlClientEncryptionKeyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CosmosDBSqlClientEncryptionKeyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
