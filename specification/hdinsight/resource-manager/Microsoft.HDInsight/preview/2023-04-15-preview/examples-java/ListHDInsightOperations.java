/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/hdinsight/resource-manager/Microsoft.HDInsight/preview/2023-04-15-preview/examples/ListHDInsightOperations.json
     */
    /**
     * Sample code: Lists all of the available operations.
     *
     * @param manager Entry point to HDInsightManager.
     */
    public static void listsAllOfTheAvailableOperations(com.azure.resourcemanager.hdinsight.HDInsightManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
