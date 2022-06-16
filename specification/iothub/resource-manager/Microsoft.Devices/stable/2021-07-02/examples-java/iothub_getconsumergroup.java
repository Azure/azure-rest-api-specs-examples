import com.azure.core.util.Context;

/** Samples for IotHubResource GetEventHubConsumerGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_getconsumergroup.json
     */
    /**
     * Sample code: IotHubResource_ListEventHubConsumerGroups.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceListEventHubConsumerGroups(
        com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .getEventHubConsumerGroupWithResponse("myResourceGroup", "testHub", "events", "test", Context.NONE);
    }
}
