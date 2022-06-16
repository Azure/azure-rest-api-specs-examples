import com.azure.core.util.Context;

/** Samples for VirtualMachines ListHosts. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetClusterVirtualMachines.json
     */
    /**
     * Sample code: Get All hosts in the cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getAllHostsInTheCluster(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.virtualMachines().listHostsWithResponse("rg1", "cluster1", Context.NONE);
    }
}
