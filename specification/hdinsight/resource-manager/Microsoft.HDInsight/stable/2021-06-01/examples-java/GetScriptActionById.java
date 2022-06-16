import com.azure.core.util.Context;

/** Samples for ScriptActions GetExecutionDetail. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/GetScriptActionById.json
     */
    /**
     * Sample code: Get script execution history by script id.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getScriptExecutionHistoryByScriptId(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.scriptActions().getExecutionDetailWithResponse("rg1", "cluster1", "391145124054712", Context.NONE);
    }
}
