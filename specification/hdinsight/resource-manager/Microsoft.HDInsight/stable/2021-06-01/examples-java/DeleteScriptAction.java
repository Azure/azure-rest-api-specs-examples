import com.azure.core.util.Context;

/** Samples for ScriptActions Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/stable/2021-06-01/examples/DeleteScriptAction.json
     */
    /**
     * Sample code: Delete a script action on HDInsight cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void deleteAScriptActionOnHDInsightCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.scriptActions().deleteWithResponse("rg1", "cluster1", "scriptName", Context.NONE);
    }
}
