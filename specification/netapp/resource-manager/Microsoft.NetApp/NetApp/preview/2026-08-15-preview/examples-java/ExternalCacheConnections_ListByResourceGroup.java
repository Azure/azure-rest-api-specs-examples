
/**
 * Samples for ExternalCacheConnections ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalCacheConnections_ListByResourceGroup.json
     */
    /**
     * Sample code: ExternalCacheConnections_ListByResourceGroup.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void
        externalCacheConnectionsListByResourceGroup(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalCacheConnections().listByResourceGroup("contosoNetAppRg", com.azure.core.util.Context.NONE);
    }
}
