/** Samples for Devices GetNetworkSettings. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/NetworkSettingsGet.json
     */
    /**
     * Sample code: NetworkSettingsGet.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void networkSettingsGet(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .devices()
            .getNetworkSettingsWithResponse(
                "testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
