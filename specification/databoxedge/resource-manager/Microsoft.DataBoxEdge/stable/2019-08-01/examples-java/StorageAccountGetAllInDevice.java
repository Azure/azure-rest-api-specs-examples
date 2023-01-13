/** Samples for StorageAccounts ListByDataBoxEdgeDevice. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/StorageAccountGetAllInDevice.json
     */
    /**
     * Sample code: StorageAccountGetAllInDevice.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void storageAccountGetAllInDevice(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .storageAccounts()
            .listByDataBoxEdgeDevice("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
