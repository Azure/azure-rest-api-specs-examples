using Azure;
using Azure.ResourceManager;
using System;
using System.Threading.Tasks;
using Azure.Core;
using Azure.Identity;
using Azure.ResourceManager.EventHubs.Models;
using Azure.ResourceManager.EventHubs;

// Generated from example definition: specification/eventhub/resource-manager/Microsoft.EventHub/Eventhub/preview/2025-05-01-preview/examples/EventHubs/EHEventHubWithCompactPolicyCreate.json
// this example is just showing the usage of "EventHubs_CreateOrUpdate" operation, for the dependent resources, they will have to be created separately.

// get your azure access token, for more details of how Azure SDK get your access token, please refer to https://learn.microsoft.com/en-us/dotnet/azure/sdk/authentication?tabs=command-line
TokenCredential cred = new DefaultAzureCredential();
// authenticate your client
ArmClient client = new ArmClient(cred);

// this example assumes you already have this EventHubResource created on azure
// for more information of creating EventHubResource, please refer to the document of EventHubResource
string subscriptionId = "5f750a97-50d9-4e36-8081-c9ee4c0210d4";
string resourceGroupName = "Default-NotificationHubs-AustraliaEast";
string namespaceName = "sdk-Namespace-5357";
string eventHubName = "sdk-EventHub-6547";
ResourceIdentifier eventHubResourceId = EventHubResource.CreateResourceIdentifier(subscriptionId, resourceGroupName, namespaceName, eventHubName);
EventHubResource eventHub = client.GetEventHubResource(eventHubResourceId);

// invoke the operation
EventHubData data = new EventHubData
{
    PartitionCount = 4L,
    Status = EventHubEntityStatus.Active,
    CaptureDescription = new CaptureDescription
    {
        Enabled = true,
        Encoding = EncodingCaptureDescription.Avro,
        IntervalInSeconds = 120,
        SizeLimitInBytes = 10485763,
        Destination = new EventHubDestination
        {
            Name = "EventHubArchive.AzureBlockBlob",
            Identity = new EventHubsCaptureIdentity
            {
                IdentityType = EventHubsCaptureIdentityType.UserAssigned,
                UserAssignedIdentity = "/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2",
            },
            StorageAccountResourceId = new ResourceIdentifier("/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage"),
            BlobContainer = "container",
            ArchiveNameFormat = "{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}",
        },
    },
    RetentionDescription = new RetentionDescription
    {
        CleanupPolicy = CleanupPolicyRetentionDescription.Compaction,
        MinCompactionLagTimeInMinutes = 10L,
        TombstoneRetentionTimeInHours = 1,
    },
    MessageTimestampType = EventHubsTimestampType.LogAppend,
    UserMetadata = "key",
};
ArmOperation<EventHubResource> lro = await eventHub.UpdateAsync(WaitUntil.Completed, data);
EventHubResource result = lro.Value;

// the variable result is a resource, you could call other operations on this instance as well
// but just for demo, we get its data from this resource instance
EventHubData resourceData = result.Data;
// for demo we just print out the id
Console.WriteLine($"Succeeded on id: {resourceData.Id}");
