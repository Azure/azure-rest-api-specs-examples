/** Samples for SqlPoolOperationResults GetLocationHeaderResult. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetLocationHeaderResultWithSqlPool.json
     */
    /**
     * Sample code: Get the result of an operation on a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getTheResultOfAnOperationOnASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolOperationResults()
            .getLocationHeaderResult(
                "ExampleResourceGroup",
                "ExampleWorkspace",
                "ExampleSqlPool",
                "fedcba98-7654-4210-fedc-ba9876543210",
                com.azure.core.util.Context.NONE);
    }
}
