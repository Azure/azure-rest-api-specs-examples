/** Samples for Shares Refresh. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/ShareRefreshPost.json
     */
    /**
     * Sample code: ShareRefreshPost.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void shareRefreshPost(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .shares()
            .refresh("testedgedevice", "smbshare", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
