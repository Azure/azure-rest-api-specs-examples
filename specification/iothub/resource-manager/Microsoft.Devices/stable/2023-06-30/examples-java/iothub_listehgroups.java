/** Samples for IotHubResource ListEventHubConsumerGroups. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_listehgroups.json
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
            .listEventHubConsumerGroups("myResourceGroup", "testHub", "events", com.azure.core.util.Context.NONE);
    }
}
