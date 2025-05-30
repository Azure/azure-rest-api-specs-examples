using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Sql.Models;
using Azure.ResourceManager.Sql;

// Generated from example definition: specification/sql/resource-manager/Microsoft.Sql/preview/2022-08-01-preview/examples/ManagedLedgerDigestUploadsGet.json
// this example is just showing the usage of "ManagedLedgerDigestUploads_Get" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this ManagedDatabaseResource created on azure
// for more information of creating ManagedDatabaseResource, please refer to the document of ManagedDatabaseResource
string subscriptionId = "00000000-1111-2222-3333-444444444444";
string resourceGroupName = "ledgertestrg";
string managedInstanceName = "ledgertestserver";
string databaseName = "testdb";
ResourceIdentifier managedDatabaseResourceId = ManagedDatabaseResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, managedInstanceName, databaseName);
ManagedDatabaseResource managedDatabase = client.GetManagedDatabaseResource(managedDatabaseResourceId);

// get the collection of this ManagedLedgerDigestUploadResource
ManagedLedgerDigestUploadCollection collection = managedDatabase.GetManagedLedgerDigestUploads();

// invoke the operation
ManagedLedgerDigestUploadsName ledgerDigestUploads = ManagedLedgerDigestUploadsName.Current;
NullableResponse<ManagedLedgerDigestUploadResource> response = await collection.GetIfExistsAsync(ledgerDigestUploads);
ManagedLedgerDigestUploadResource result = response.HasValue ? response.Value : null;

if (result == null)
{
    Console.WriteLine("Succeeded with null as result");
}
else
{
    // the variable result is a resource, you could call other operations on this instance as well
    // but just for demo, we get its data from this resource instance
    ManagedLedgerDigestUploadData resourceData = result.Data;
    // for demo we just print out the id
    Console.WriteLine($"Succeeded on id: {resourceData.Id}");
}
