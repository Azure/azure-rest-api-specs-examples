import com.azure.core.util.Context;

/** Samples for ManagedClusters GetUpgradeProfile. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-02-01/examples/ManagedClustersGetUpgradeProfile.json
     */
    /**
     * Sample code: Get Upgrade Profile for Managed Cluster.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getUpgradeProfileForManagedCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .getUpgradeProfileWithResponse("rg1", "clustername1", Context.NONE);
    }
}
