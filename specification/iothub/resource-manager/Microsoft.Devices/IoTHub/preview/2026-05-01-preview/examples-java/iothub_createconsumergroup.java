
import com.azure.resourcemanager.iothub.models.EventHubConsumerGroupName;

/**
 * Samples for IotHubResource CreateEventHubConsumerGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_createconsumergroup.json
     */
    /**
     * Sample code: IotHubResource_CreateEventHubConsumerGroup.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void
        iotHubResourceCreateEventHubConsumerGroup(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().defineEventHubConsumerGroup("test")
            .withExistingEventHubEndpoint("myResourceGroup", "testHub", "events")
            .withProperties(new EventHubConsumerGroupName().withName("test")).create();
    }
}
