/** Samples for StorageAccounts Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/StorageAccountDelete.json
     */
    /**
     * Sample code: StorageAccountDelete.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void storageAccountDelete(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .storageAccounts()
            .delete("testedgedevice", "storageaccount1", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
