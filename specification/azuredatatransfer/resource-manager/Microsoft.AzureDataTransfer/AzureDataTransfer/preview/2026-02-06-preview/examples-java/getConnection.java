
/**
 * Samples for Connections GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-06-preview/getConnection.json
     */
    /**
     * Sample code: Gets connection resource.
     * 
     * @param manager Entry point to DataTransferManager.
     */
    public static void getsConnectionResource(com.azure.resourcemanager.datatransfer.DataTransferManager manager) {
        manager.connections().getByResourceGroupWithResponse("testRG", "testConnection",
            com.azure.core.util.Context.NONE);
    }
}
