import com.azure.core.util.Context;

/** Samples for SqlPoolMetadataSyncConfigs Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolMetadataSyncConfig.json
     */
    /**
     * Sample code: Get metadata sync config for a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getMetadataSyncConfigForASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolMetadataSyncConfigs()
            .getWithResponse("ExampleResourceGroup", "ExampleWorkspace", "ExampleSqlPool", Context.NONE);
    }
}
