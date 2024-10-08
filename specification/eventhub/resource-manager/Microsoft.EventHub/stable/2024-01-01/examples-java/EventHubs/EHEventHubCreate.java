
import com.azure.resourcemanager.eventhubs.fluent.models.EventhubInner;
import com.azure.resourcemanager.eventhubs.models.CaptureDescription;
import com.azure.resourcemanager.eventhubs.models.CaptureIdentity;
import com.azure.resourcemanager.eventhubs.models.CaptureIdentityType;
import com.azure.resourcemanager.eventhubs.models.CleanupPolicyRetentionDescription;
import com.azure.resourcemanager.eventhubs.models.Destination;
import com.azure.resourcemanager.eventhubs.models.EncodingCaptureDescription;
import com.azure.resourcemanager.eventhubs.models.EntityStatus;
import com.azure.resourcemanager.eventhubs.models.RetentionDescription;

/**
 * Samples for EventHubs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/eventhub/resource-manager/Microsoft.EventHub/stable/2024-01-01/examples/EventHubs/EHEventHubCreate.
     * json
     */
    /**
     * Sample code: EventHubCreate.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.eventHubs().manager().serviceClient().getEventHubs().createOrUpdateWithResponse(
            "Default-NotificationHubs-AustraliaEast", "sdk-Namespace-5357", "sdk-EventHub-6547",
            new EventhubInner().withMessageRetentionInDays(4L).withPartitionCount(4L).withStatus(EntityStatus.ACTIVE)
                .withUserMetadata("key")
                .withCaptureDescription(new CaptureDescription().withEnabled(true)
                    .withEncoding(EncodingCaptureDescription.AVRO).withIntervalInSeconds(120)
                    .withSizeLimitInBytes(10485763)
                    .withDestination(new Destination().withName("EventHubArchive.AzureBlockBlob").withIdentity(
                        new CaptureIdentity().withType(CaptureIdentityType.USER_ASSIGNED).withUserAssignedIdentity(
                            "/subscriptions/SampleSubscription/resourceGroups/ResurceGroupSample/providers/Microsoft.ManagedIdentity/userAssignedIdentities/ud2"))
                        .withStorageAccountResourceId(
                            "/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage")
                        .withBlobContainer("container").withArchiveNameFormat(
                            "{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}")))
                .withRetentionDescription(
                    new RetentionDescription().withCleanupPolicy(CleanupPolicyRetentionDescription.COMPACT)
                        .withRetentionTimeInHours(96L).withTombstoneRetentionTimeInHours(1)),
            com.azure.core.util.Context.NONE);
    }
}
