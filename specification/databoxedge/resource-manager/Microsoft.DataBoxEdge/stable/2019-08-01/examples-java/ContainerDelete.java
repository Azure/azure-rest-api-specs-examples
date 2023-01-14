/** Samples for Containers Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/ContainerDelete.json
     */
    /**
     * Sample code: ContainerDelete.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void containerDelete(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .containers()
            .delete(
                "testedgedevice",
                "storageaccount1",
                "blobcontainer1",
                "GroupForEdgeAutomation",
                com.azure.core.util.Context.NONE);
    }
}
