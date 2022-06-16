import com.azure.core.util.Context;

/** Samples for ScriptExecutionHistory ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetScriptExecutionHistory.json
     */
    /**
     * Sample code: Get Script Execution History List.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getScriptExecutionHistoryList(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.scriptExecutionHistories().listByCluster("rg1", "cluster1", Context.NONE);
    }
}
