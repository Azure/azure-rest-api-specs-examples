
/**
 * Samples for IotHubResource ListEventHubConsumerGroups.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/iothub_listehgroups.json
     */
    /**
     * Sample code: IotHubResource_ListEventHubConsumerGroups.
     * 
     * @param manager Entry point to IotHubManager.
     */
    public static void
        iotHubResourceListEventHubConsumerGroups(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().listEventHubConsumerGroups("myResourceGroup", "testHub", "events",
            com.azure.core.util.Context.NONE);
    }
}
