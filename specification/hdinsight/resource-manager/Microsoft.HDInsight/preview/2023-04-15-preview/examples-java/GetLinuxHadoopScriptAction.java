/** Samples for ScriptActions ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/GetLinuxHadoopScriptAction.json
     */
    /**
     * Sample code: List all persisted script actions for the given cluster.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void listAllPersistedScriptActionsForTheGivenCluster(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.scriptActions().listByCluster("rg1", "cluster1", com.azure.core.util.Context.NONE);
    }
}
