import com.azure.core.util.Context;

/** Samples for ManagedClusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2021-10-01/examples/ManagedClustersList.json
     */
    /**
     * Sample code: List Managed Clusters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedClusters(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters().list(Context.NONE);
    }
}
