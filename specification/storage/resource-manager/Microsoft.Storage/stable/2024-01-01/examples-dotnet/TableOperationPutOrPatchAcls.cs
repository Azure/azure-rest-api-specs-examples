using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2024-01-01/examples/TableOperationPutOrPatchAcls.json
// this example is just showing the usage of "Table_Update" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this TableResource created on azure
// for more information of creating TableResource, please refer to the document of TableResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res3376";
string accountName = "sto328";
string tableName = "table6185";
ResourceIdentifier tableResourceId = TableResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, tableName);
TableResource table = client.GetTableResource(tableResourceId);

// invoke the operation
TableData data = new TableData
{
    SignedIdentifiers = {new StorageTableSignedIdentifier("MTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODkwMTI")
    {
    AccessPolicy = new StorageTableAccessPolicy("raud")
    {
    StartOn = DateTimeOffset.Parse("2022-03-17T08:49:37.0000000Z"),
    ExpireOn = DateTimeOffset.Parse("2022-03-20T08:49:37.0000000Z"),
    },
    }, new StorageTableSignedIdentifier("PTIzNDU2Nzg5MDEyMzQ1Njc4OTAxMjM0NTY3ODklMTI")
    {
    AccessPolicy = new StorageTableAccessPolicy("rad")
    {
    StartOn = DateTimeOffset.Parse("2022-03-17T08:49:37.0000000Z"),
    ExpireOn = DateTimeOffset.Parse("2022-03-20T08:49:37.0000000Z"),
    },
    }},
};
TableResource result = await table.UpdateAsync(data);

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
TableData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
