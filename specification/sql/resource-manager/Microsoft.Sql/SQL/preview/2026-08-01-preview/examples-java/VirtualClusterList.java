
/**
 * Samples for VirtualClusters List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-01-preview/VirtualClusterList.json
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
