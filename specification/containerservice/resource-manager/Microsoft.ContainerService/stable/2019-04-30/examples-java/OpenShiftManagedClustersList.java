import com.azure.core.util.Context;

/** Samples for OpenShiftManagedClusters List. */
public final class Main {
    /*
     * x-ms-original-file: specification/containerservice/resource-manager/Microsoft.ContainerService/stable/2019-04-30/examples/OpenShiftManagedClustersList.json
     */
    /**
     * Sample code: List Managed Clusters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedClusters(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getOpenShiftManagedClusters().list(Context.NONE);
    }
}
