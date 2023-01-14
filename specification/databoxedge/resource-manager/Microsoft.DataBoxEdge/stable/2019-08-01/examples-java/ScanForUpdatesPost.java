/** Samples for Devices ScanForUpdates. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/ScanForUpdatesPost.json
     */
    /**
     * Sample code: ScanForUpdatesPost.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void scanForUpdatesPost(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager.devices().scanForUpdates("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
