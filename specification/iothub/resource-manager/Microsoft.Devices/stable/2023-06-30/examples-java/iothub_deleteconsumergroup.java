/** Samples for IotHubResource DeleteEventHubConsumerGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_deleteconsumergroup.json
     */
    /**
     * Sample code: IotHubResource_DeleteEventHubConsumerGroup.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceDeleteEventHubConsumerGroup(
        com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager
            .iotHubResources()
            .deleteEventHubConsumerGroupWithResponse(
                "myResourceGroup", "testHub", "events", "test", com.azure.core.util.Context.NONE);
    }
}
