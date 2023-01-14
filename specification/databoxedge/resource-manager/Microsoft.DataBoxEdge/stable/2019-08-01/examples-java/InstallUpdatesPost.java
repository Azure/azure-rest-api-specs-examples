/** Samples for Devices InstallUpdates. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/InstallUpdatesPost.json
     */
    /**
     * Sample code: InstallUpdatesPost.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void installUpdatesPost(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager.devices().installUpdates("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
