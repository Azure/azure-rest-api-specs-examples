import com.azure.core.util.Context;
import java.util.Arrays;

/** Samples for VirtualMachines RestartHosts. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/RestartVirtualMachinesOperation.json
     */
    /**
     * Sample code: Restarts the specified HDInsight cluster hosts.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void restartsTheSpecifiedHDInsightClusterHosts(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.virtualMachines().restartHosts("rg1", "cluster1", Arrays.asList("gateway1", "gateway3"), Context.NONE);
    }
}
