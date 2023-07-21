/** Samples for ScriptActions GetExecutionAsyncOperationStatus. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/GetScriptExecutionAsyncOperationStatus.json
     */
    /**
     * Sample code: Gets the async execution operation status.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void getsTheAsyncExecutionOperationStatus(
        com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager
            .scriptActions()
            .getExecutionAsyncOperationStatusWithResponse(
                "rg1", "cluster1", "CF938302-6B4D-44A0-A6D2-C0D67E847AEC", com.azure.core.util.Context.NONE);
    }
}
