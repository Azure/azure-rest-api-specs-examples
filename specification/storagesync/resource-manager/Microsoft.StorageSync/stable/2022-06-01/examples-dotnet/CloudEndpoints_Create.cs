using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.StorageSync.Models;
using Azure.ResourceManager.StorageSync;

// Generated from example definition: specification/storagesync/resource-manager/Microsoft.StorageSync/stable/2022-06-01/examples/CloudEndpoints_Create.json
// this example is just showing the usage of "CloudEndpoints_Create" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this CloudEndpointResource created on azure
// for more information of creating CloudEndpointResource, please refer to the document of CloudEndpointResource
string subscriptionId = "52b8da2f-61e0-4a1f-8dde-336911f367fb";
string resourceGroupName = "SampleResourceGroup_1";
string storageSyncServiceName = "SampleStorageSyncService_1";
string syncGroupName = "SampleSyncGroup_1";
string cloudEndpointName = "SampleCloudEndpoint_1";
ResourceIdentifier cloudEndpointResourceId = CloudEndpointResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, storageSyncServiceName, syncGroupName, cloudEndpointName);
CloudEndpointResource cloudEndpoint = client.GetCloudEndpointResource(cloudEndpointResourceId);

// invoke the operation
CloudEndpointCreateOrUpdateContent content = new CloudEndpointCreateOrUpdateContent
{
    StorageAccountResourceId = new ResourceIdentifier("/subscriptions/744f4d70-6d17-4921-8970-a765d14f763f/resourceGroups/tminienv59svc/providers/Microsoft.Storage/storageAccounts/tminienv59storage"),
    AzureFileShareName = "cvcloud-afscv-0719-058-a94a1354-a1fd-4e9a-9a50-919fad8c4ba4",
    StorageAccountTenantId = Guid.Parse("\"72f988bf-86f1-41af-91ab-2d7cd011db47\""),
    FriendlyName = "ankushbsubscriptionmgmtmab",
};
ArmOperation<CloudEndpointResource> lro = await cloudEndpoint.UpdateAsync(WaitUntil.Completed, content);
CloudEndpointResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
CloudEndpointData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
