using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Storage.Models;
using Azure.ResourceManager.Storage;

// Generated from example definition: specification/storage/resource-manager/Microsoft.Storage/stable/2023-05-01/examples/StorageAccountSetBlobInventoryPolicyIncludeDeleteAndNewSchemaForNonHnsAccount.json
// this example is just showing the usage of "BlobInventoryPolicies_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this BlobInventoryPolicyResource created on azure
// for more information of creating BlobInventoryPolicyResource, please refer to the document of BlobInventoryPolicyResource
string subscriptionId = "{subscription-id}";
string resourceGroupName = "res7687";
string accountName = "sto9699";
BlobInventoryPolicyName blobInventoryPolicyName = BlobInventoryPolicyName.Default;
ResourceIdentifier blobInventoryPolicyResourceId = BlobInventoryPolicyResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, accountName, blobInventoryPolicyName);
BlobInventoryPolicyResource blobInventoryPolicy = client.GetBlobInventoryPolicyResource(blobInventoryPolicyResourceId);

// invoke the operation
BlobInventoryPolicyData data = new BlobInventoryPolicyData()
{
    PolicySchema = new BlobInventoryPolicySchema(true, BlobInventoryRuleType.Inventory, new BlobInventoryPolicyRule[]
{
new BlobInventoryPolicyRule(true,"inventoryPolicyRule1","container1",new BlobInventoryPolicyDefinition(BlobInventoryPolicyFormat.Csv,BlobInventoryPolicySchedule.Daily,BlobInventoryPolicyObjectType.Blob,new string[]
{
"Name","Creation-Time","Last-Modified","Content-Length","Content-MD5","BlobType","AccessTier","AccessTierChangeTime","Snapshot","VersionId","IsCurrentVersion","Tags","ContentType","ContentEncoding","ContentLanguage","ContentCRC64","CacheControl","Metadata","Deleted","RemainingRetentionDays"
})
{
Filters = new BlobInventoryPolicyFilter()
{
IncludePrefix =
{
"inventoryprefix1","inventoryprefix2"
},
ExcludePrefix =
{
"excludeprefix1","excludeprefix2"
},
BlobTypes =
{
"blockBlob","appendBlob","pageBlob"
},
IncludeBlobVersions = true,
IncludeSnapshots = true,
IncludeDeleted = true,
},
}),new BlobInventoryPolicyRule(true,"inventoryPolicyRule2","container2",new BlobInventoryPolicyDefinition(BlobInventoryPolicyFormat.Parquet,BlobInventoryPolicySchedule.Weekly,BlobInventoryPolicyObjectType.Container,new string[]
{
"Name","Last-Modified","Metadata","LeaseStatus","LeaseState","LeaseDuration","PublicAccess","HasImmutabilityPolicy","HasLegalHold","Etag","DefaultEncryptionScope","DenyEncryptionScopeOverride","ImmutableStorageWithVersioningEnabled","Deleted","Version","DeletedTime","RemainingRetentionDays"
}))
}),
};
ArmOperation<BlobInventoryPolicyResource> lro = await blobInventoryPolicy.CreateOrUpdateAsync(WaitUntil.Completed, data);
BlobInventoryPolicyResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
BlobInventoryPolicyData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
