import com.azure.core.util.Context;

/** Samples for VirtualMachines GetAsyncOperationStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetRestartHostsAsyncOperationStatus.json
     */
    /**
     * Sample code: Gets the async operation status of restarting host.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getsTheAsyncOperationStatusOfRestartingHost(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .virtualMachines()
            .getAsyncOperationStatusWithResponse(
                "rg1", "cluster1", "CF938302-6B4D-44A0-A6D2-C0D67E847AEC", Context.NONE);
    }
}
