
/**
 * Samples for Connections Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-01/Connections_Get.json
     */
    /**
     * Sample code: Connections_Get.
     * 
     * @param manager Entry point to StorageMoverManager.
     */
    public static void connectionsGet(com.azure.resourcemanager.storagemover.StorageMoverManager manager) {
        manager.connections().getWithResponse("examples-rg", "examples-storageMoverName", "example-connection",
            com.azure.core.util.Context.NONE);
    }
}
