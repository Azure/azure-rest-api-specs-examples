import com.azure.core.util.Context;

/** Samples for ManagedClusters ListOutboundNetworkDependenciesEndpoints. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2022-07-01/examples/OutboundNetworkDependenciesEndpointsList.json
     */
    /**
     * Sample code: List OutboundNetworkDependenciesEndpoints by Managed Cluster.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listOutboundNetworkDependenciesEndpointsByManagedCluster(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .kubernetesClusters()
            .manager()
            .serviceClient()
            .getManagedClusters()
            .listOutboundNetworkDependenciesEndpoints("rg1", "clustername1", Context.NONE);
    }
}
