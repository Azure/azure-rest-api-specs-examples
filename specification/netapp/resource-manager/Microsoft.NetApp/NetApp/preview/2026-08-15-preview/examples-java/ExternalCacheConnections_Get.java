
/**
 * Samples for ExternalCacheConnections GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalCacheConnections_Get.json
     */
    /**
     * Sample code: ExternalCacheConnections_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void externalCacheConnectionsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalCacheConnections().getByResourceGroupWithResponse("contosoNetAppRg", "contosoCacheConnection",
            com.azure.core.util.Context.NONE);
    }
}
