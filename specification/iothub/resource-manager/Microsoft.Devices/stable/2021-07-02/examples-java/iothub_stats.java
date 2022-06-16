import com.azure.core.util.Context;

/** Samples for IotHubResource GetStats. */
public final class Main {
    /*
     * x-ms-original-file: specification/iothub/resource-manager/Microsoft.Devices/stable/2021-07-02/examples/iothub_stats.json
     */
    /**
     * Sample code: IotHubResource_GetStats.
     *
     * @param manager Entry point to IotHubManager.
     */
    public static void iotHubResourceGetStats(com.azure.resourcemanager.iothub.IotHubManager manager) {
        manager.iotHubResources().getStatsWithResponse("myResourceGroup", "testHub", Context.NONE);
    }
}
