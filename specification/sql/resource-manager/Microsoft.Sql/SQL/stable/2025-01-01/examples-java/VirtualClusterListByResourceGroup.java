
/**
 * Samples for VirtualClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/VirtualClusterListByResourceGroup.json
     */
    /**
     * Sample code: List virtual clusters by resource group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listVirtualClustersByResourceGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getVirtualClusters().listByResourceGroup("testrg", com.azure.core.util.Context.NONE);
    }
}
