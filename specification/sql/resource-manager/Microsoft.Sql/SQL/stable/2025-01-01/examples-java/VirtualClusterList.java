
/**
 * Samples for VirtualClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/VirtualClusterList.json
     */
    /**
     * Sample code: List virtualClusters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listVirtualClusters(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getVirtualClusters().list(com.azure.core.util.Context.NONE);
    }
}
