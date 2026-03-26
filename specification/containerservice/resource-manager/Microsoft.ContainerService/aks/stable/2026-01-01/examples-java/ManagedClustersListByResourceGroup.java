
/**
 * Samples for ManagedClusters ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-01-01/ManagedClustersListByResourceGroup.json
     */
    /**
     * Sample code: Get Managed Clusters by Resource Group.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getManagedClustersByResourceGroup(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().listByResourceGroup("rg1", com.azure.core.util.Context.NONE);
    }
}
