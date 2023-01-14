/** Samples for Containers Refresh. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/ContainerRefresh.json
     */
    /**
     * Sample code: ContainerRefresh.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void containerRefresh(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .containers()
            .refresh(
                "testedgedevice",
                "storageaccount1",
                "blobcontainer1",
                "GroupForEdgeAutomation",
                com.azure.core.util.Context.NONE);
    }
}
