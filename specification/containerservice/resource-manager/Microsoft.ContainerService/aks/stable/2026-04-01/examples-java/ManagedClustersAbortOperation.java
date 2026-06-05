
/**
 * Samples for ManagedClusters AbortLatestOperation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01/ManagedClustersAbortOperation.json
     */
    /**
     * Sample code: Abort operation on managed cluster.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        abortOperationOnManagedCluster(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().abortLatestOperation("rg1", "clustername1",
            com.azure.core.util.Context.NONE);
    }
}
