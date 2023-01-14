/** Samples for Shares Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/ShareGet.json
     */
    /**
     * Sample code: ShareGet.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void shareGet(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .shares()
            .getWithResponse("testedgedevice", "smbshare", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
