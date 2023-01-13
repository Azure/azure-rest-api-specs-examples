/** Samples for StorageAccounts Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/StorageAccountGet.json
     */
    /**
     * Sample code: StorageAccountGet.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void storageAccountGet(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .storageAccounts()
            .getWithResponse(
                "testedgedevice", "blobstorageaccount1", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
