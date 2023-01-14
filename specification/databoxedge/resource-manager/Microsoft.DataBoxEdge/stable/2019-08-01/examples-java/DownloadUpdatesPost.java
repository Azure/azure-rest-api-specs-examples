/** Samples for Devices DownloadUpdates. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/DownloadUpdatesPost.json
     */
    /**
     * Sample code: DownloadUpdatesPost.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void downloadUpdatesPost(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager.devices().downloadUpdates("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
