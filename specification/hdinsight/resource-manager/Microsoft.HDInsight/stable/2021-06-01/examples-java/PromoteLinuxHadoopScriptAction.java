import com.azure.core.util.Context;

/** Samples for ScriptExecutionHistory Promote. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/PromoteLinuxHadoopScriptAction.json
     */
    /**
     * Sample code: Promote a script action on HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void promoteAScriptActionOnHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.scriptExecutionHistories().promoteWithResponse("rg1", "cluster1", "391145124054712", Context.NONE);
    }
}
