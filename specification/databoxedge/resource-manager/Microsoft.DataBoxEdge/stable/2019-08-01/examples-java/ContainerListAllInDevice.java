/** Samples for Containers ListByStorageAccount. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/ContainerListAllInDevice.json
     */
    /**
     * Sample code: ContainerListAllInDevice.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void containerListAllInDevice(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .containers()
            .listByStorageAccount(
                "testedgedevice", "storageaccount1", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
