import com.azure.core.util.Context;

/** Samples for ManagedClusters ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-10-01/examples/ManagedClustersListByResourceGroup.json
     */
    /**
     * Sample code: Get Managed Clusters by Resource Group.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getManagedClustersByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .listByResourceGroup("rg1", Context.NONE);
    }
}
