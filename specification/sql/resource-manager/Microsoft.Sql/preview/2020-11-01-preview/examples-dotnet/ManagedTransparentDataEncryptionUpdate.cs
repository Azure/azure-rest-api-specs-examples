using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2020-11-01-preview/examples/ManagedTransparentDataEncryptionUpdate.json
// this example is just showing the usage of "ManagedDatabaseTransparentDataEncryption_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDatabaseResource created on azure
// for more information of creating ManagedDatabaseResource, please refer to the document of ManagedDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "securitytde-42-rg";
string managedInstanceName = "securitytde-42";
string databaseName = "testdb";
ResourceIdentifier managedDatabaseResourceId = ManagedDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName);
ManagedDatabaseResource managedDatabase = client.GetManagedDatabaseResource(managedDatabaseResourceId);

// get the collection of this ManagedTransparentDataEncryptionResource
ManagedTransparentDataEncryptionCollection collection = managedDatabase.GetManagedTransparentDataEncryptions();

// invoke the operation
TransparentDataEncryptionName tdeName = TransparentDataEncryptionName.Current;
ManagedTransparentDataEncryptionData data = new ManagedTransparentDataEncryptionData()
{
    State = TransparentDataEncryptionState.Enabled,
};
ArmOperation<ManagedTransparentDataEncryptionResource> lro = await collection.CreateOrUpdateAsync(WaitUntil.Completed, tdeName, data);
ManagedTransparentDataEncryptionResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
ManagedTransparentDataEncryptionData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
