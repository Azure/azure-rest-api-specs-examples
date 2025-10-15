
/**
 * Samples for ScriptActions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2025-01-15-preview/examples/
     * DeleteScriptAction.json
     */
    /**
     * Sample code: Delete a script action on HDInsight cluster.
     * 
     * @param manager Entry point to HDInsightManager.
     */
    public static void
        deleteAScriptActionOnHDInsightCluster(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.scriptActions().deleteWithResponse("rg1", "cluster1", "scriptName", com.azure.core.util.Context.NONE);
    }
}
