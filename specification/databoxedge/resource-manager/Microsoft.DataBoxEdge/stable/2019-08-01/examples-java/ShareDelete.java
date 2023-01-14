/** Samples for Shares Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/ShareDelete.json
     */
    /**
     * Sample code: ShareDelete.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void shareDelete(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .shares()
            .delete("testedgedevice", "smbshare", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
