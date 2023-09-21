/** Samples for IotHubResource GetStats. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2023-06-30/examples/iothub_stats.json
     */
    /**
     * Sample code: IotHubResource_GetStats.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetStats(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getStatsWithResponse("myResourceGroup", "testHub", com.azure.core.util.Context.NONE);
    }
}
