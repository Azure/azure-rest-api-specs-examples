using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.Resources;
using Azure.ResourceManager.StorageActions.Models;
using Azure.ResourceManager.StorageActions;

// Generated from example definition: 2023-01-01/misc/PerformStorageTaskActionsPreview.json
// this example is just showing the usage of "StorageTasksOperationGroup_PreviewActions" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this SubscriptionResource created on azure
// for more information of creating SubscriptionResource, please refer to the document of SubscriptionResource
string subscriptionId = "1f31ba14-ce16-4281-b9b4-3e78da6e1616";
ResourceIdentifier subscriptionResourceId = SubscriptionResource.CreateResourceIdentifier(subscriptionId);
SubscriptionResource subscriptionResource = client.GetSubscriptionResource(subscriptionResourceId);

// invoke the operation
AzureLocation location = new AzureLocation("eastus");
StorageTaskPreviewAction storageTaskPreviewAction = new StorageTaskPreviewAction(new StorageTaskPreviewActionProperties(new StorageTaskPreviewContainerProperties
{
    Name = "firstContainer",
    Metadata = {new StorageTaskPreviewKeyValueProperties
    {
    Key = "mContainerKey1",
    Value = "mContainerValue1",
    }},
}, new StorageTaskPreviewBlobProperties[]
{
new StorageTaskPreviewBlobProperties
{
Name = "folder1/file1.txt",
Properties = {new StorageTaskPreviewKeyValueProperties
{
Key = "Creation-Time",
Value = "Wed, 07 Jun 2023 05:23:29 GMT",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Last-Modified",
Value = "Wed, 07 Jun 2023 05:23:29 GMT",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Etag",
Value = "0x8DB67175454D36D",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Content-Length",
Value = "38619",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Content-Type",
Value = "text/xml",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Content-Encoding",
Value = "",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Content-Language",
Value = "",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Content-CRC64",
Value = "",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Content-MD5",
Value = "njr6iDrmU9+FC89WMK22EA==",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Cache-Control",
Value = "",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Content-Disposition",
Value = "",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "BlobType",
Value = "BlockBlob",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "AccessTier",
Value = "Hot",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "AccessTierInferred",
Value = "true",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "LeaseStatus",
Value = "unlocked",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "LeaseState",
Value = "available",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "ServerEncrypted",
Value = "true",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "TagCount",
Value = "1",
}},
Metadata = {new StorageTaskPreviewKeyValueProperties
{
Key = "mKey1",
Value = "mValue1",
}},
Tags = {new StorageTaskPreviewKeyValueProperties
{
Key = "tKey1",
Value = "tValue1",
}},
},
new StorageTaskPreviewBlobProperties
{
Name = "folder2/file1.txt",
Properties = {new StorageTaskPreviewKeyValueProperties
{
Key = "Creation-Time",
Value = "Wed, 06 Jun 2023 05:23:29 GMT",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Last-Modified",
Value = "Wed, 06 Jun 2023 05:23:29 GMT",
}, new StorageTaskPreviewKeyValueProperties
{
Key = "Etag",
Value = "0x6FB67175454D36D",
}},
Metadata = {new StorageTaskPreviewKeyValueProperties
{
Key = "mKey2",
Value = "mValue2",
}},
Tags = {new StorageTaskPreviewKeyValueProperties
{
Key = "tKey2",
Value = "tValue2",
}},
}
}, new StorageTaskPreviewActionCondition(new StorageTaskPreviewActionIfCondition
{
    Condition = "[[equals(AccessTier, 'Hot')]]",
}, true)));
StorageTaskPreviewAction result = await subscriptionResource.PreviewActionsAsync(location, storageTaskPreviewAction);

Console.WriteLine($"Succeeded: {result}");
