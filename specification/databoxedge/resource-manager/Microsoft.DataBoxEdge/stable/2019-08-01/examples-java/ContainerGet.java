/** Samples for Containers Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/ContainerGet.json
     */
    /**
     * Sample code: ContainerGet.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void containerGet(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .containers()
            .getWithResponse(
                "testedgedevice",
                "storageaccount1",
                "blobcontainer1",
                "GroupForEdgeAutomation",
                com.azure.core.util.Context.NONE);
    }
}
