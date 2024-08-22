
/**
 * Samples for VirtualMachines ListHosts.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2024-08-01-preview/examples/
     * GetClusterVirtualMachines.json
     */
    /**
     * Sample code: Get All hosts in the cluster.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAllHostsInTheCluster(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.virtualMachines().listHostsWithResponse("rg1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
