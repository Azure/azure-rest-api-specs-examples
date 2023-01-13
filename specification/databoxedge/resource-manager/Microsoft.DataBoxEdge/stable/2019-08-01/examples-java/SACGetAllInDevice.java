/** Samples for StorageAccountCredentials ListByDataBoxEdgeDevice. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/SACGetAllInDevice.json
     */
    /**
     * Sample code: SACGetAllInDevice.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void sACGetAllInDevice(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .storageAccountCredentials()
            .listByDataBoxEdgeDevice("testedgedevice", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
