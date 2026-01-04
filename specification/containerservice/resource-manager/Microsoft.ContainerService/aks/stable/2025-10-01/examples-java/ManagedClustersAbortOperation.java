
/**
 * Samples for ManagedClusters AbortLatestOperation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/containerservice/resource-manager/Microsoft.ContainerService/aks/stable/2025-10-01/examples/
     * ManagedClustersAbortOperation.json
     */
    /**
     * Sample code: Abort operation on managed cluster.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void abortOperationOnManagedCluster(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.kubernetesClusters().manager().serviceClient().getManagedClusters().abortLatestOperation("rg1",
            "clustername1", com.azure.core.util.Context.NONE);
    }
}
