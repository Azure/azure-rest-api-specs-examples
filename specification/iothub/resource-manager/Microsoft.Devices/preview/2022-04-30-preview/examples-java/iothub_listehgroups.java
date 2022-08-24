import com.azure.core.util.Context;

/** Samples for IotHubResource ListEventHubConsumerGroups. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/preview/2022-04-30-preview/examples/iothub_listehgroups.json
     */
    /**
     * Sample code: IotHubResource_ListEventHubConsumerGroups.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceListEventHubConsumerGroups(
        com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().listEventHubConsumerGroups("myResourceGroup", "testHub", "events", Context.NONE);
    }
}
