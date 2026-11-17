
import com.azure.resourcemanager.eventhubs.fluent.models.EventhubInner;
import com.azure.resourcemanager.eventhubs.models.CaptureDescription;
import com.azure.resourcemanager.eventhubs.models.CaptureIdentity;
import com.azure.resourcemanager.eventhubs.models.CaptureIdentityType;
import com.azure.resourcemanager.eventhubs.models.Destination;
import com.azure.resourcemanager.eventhubs.models.EncodingCaptureDescription;
import com.azure.resourcemanager.eventhubs.models.EntityStatus;

/**
 * Samples for EventHubs CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01-preview/EventHubs/EHEventHubCreate.json
     */
    /**
     * Sample code: EHEventHubCreate.
     * 
     * @param manager Entry point to EventHubsManager.
     */
    public static void eHEventHubCreate(com.azure.resourcemanager.eventhubs.EventHubsManager manager) {
        manager.serviceClient().getEventHubs().createOrUpdateWithResponse("Default-NotificationHubs-AustraliaEast",
            "sdk-Namespace-5357", "sdk-EventHub-6547",
            new EventhubInner().withPartitionCount(4L).withStatus(EntityStatus.ACTIVE)
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
                .withUserMetadata("key"),
            com.azure.core.util.Context.NONE);
    }
}
