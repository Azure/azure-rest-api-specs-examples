
/**
 * Samples for PrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/PrivateLinkResourcesList.json
     */
    /**
     * Sample code: List Private Link Resources by Managed Cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void listPrivateLinkResourcesByManagedCluster(
        com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getPrivateLinkResources().listWithResponse("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
