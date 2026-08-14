
/**
 * Samples for IotHubResource GetEventHubConsumerGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/iothub_getconsumergroup.json
     */
    /**
     * Sample code: IotHubResource_ListEventHubConsumerGroups.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void
        iotHubResourceListEventHubConsumerGroups(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getEventHubConsumerGroupWithResponse("myResourceGroup", "testHub", "events", "test",
            com.azure.core.util.Context.NONE);
    }
}
