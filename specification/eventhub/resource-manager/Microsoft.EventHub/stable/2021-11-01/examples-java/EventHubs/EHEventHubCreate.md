Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager_2.10.0/sdk/resourcemanager/azure-resourcemanager/README.md) on how to add the SDK to your project and authenticate.

```java
import com.azure.core.util.Context;
import com.azure.resourcemanager.eventhubs.fluent.models.EventhubInner;
import com.azure.resourcemanager.eventhubs.models.CaptureDescription;
import com.azure.resourcemanager.eventhubs.models.Destination;
import com.azure.resourcemanager.eventhubs.models.EncodingCaptureDescription;
import com.azure.resourcemanager.eventhubs.models.EntityStatus;

/** Samples for EventHubs CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/eventhub/resource-manager/Microsoft.EventHub/stable/2021-11-01/examples/EventHubs/EHEventHubCreate.json
     */
    /**
     * Sample code: EventHubCreate.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void eventHubCreate(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .eventHubs()
            .manager()
            .serviceClient()
            .getEventHubs()
            .createOrUpdateWithResponse(
                "Default-NotificationHubs-AustraliaEast",
                "sdk-Namespace-5357",
                "sdk-EventHub-6547",
                new EventhubInner()
                    .withMessageRetentionInDays(4L)
                    .withPartitionCount(4L)
                    .withStatus(EntityStatus.ACTIVE)
                    .withCaptureDescription(
                        new CaptureDescription()
                            .withEnabled(true)
                            .withEncoding(EncodingCaptureDescription.AVRO)
                            .withIntervalInSeconds(120)
                            .withSizeLimitInBytes(10485763)
                            .withDestination(
                                new Destination()
                                    .withName("EventHubArchive.AzureBlockBlob")
                                    .withStorageAccountResourceId(
                                        "/subscriptions/e2f361f0-3b27-4503-a9cc-21cfba380093/resourceGroups/Default-Storage-SouthCentralUS/providers/Microsoft.ClassicStorage/storageAccounts/arjunteststorage")
                                    .withBlobContainer("container")
                                    .withArchiveNameFormat(
                                        "{Namespace}/{EventHub}/{PartitionId}/{Year}/{Month}/{Day}/{Hour}/{Minute}/{Second}"))),
                Context.NONE);
    }
}
```
