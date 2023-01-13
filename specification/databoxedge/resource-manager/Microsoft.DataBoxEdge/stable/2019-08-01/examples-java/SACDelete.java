/** Samples for StorageAccountCredentials Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/databoxedge/resource-manager/Microsoft.DataBoxEdge/stable/2019-08-01/examples/SACDelete.json
     */
    /**
     * Sample code: SACDelete.
     *
     * @param manager Entry point to DataBoxEdgeManager.
     */
    public static void sACDelete(com.azure.resourcemanager.databoxedge.DataBoxEdgeManager manager) {
        manager
            .storageAccountCredentials()
            .delete("testedgedevice", "sac1", "GroupForEdgeAutomation", com.azure.core.util.Context.NONE);
    }
}
