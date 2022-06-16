import com.azure.core.util.Context;
import com.azure.resourcemanager.synapse.fluent.models.MetadataSyncConfigInner;

/** Samples for SqlPoolMetadataSyncConfigs Create. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/CreateSqlPoolMetadataSyncConfig.json
     */
    /**
     * Sample code: Set metadata sync config for a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void setMetadataSyncConfigForASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolMetadataSyncConfigs()
            .createWithResponse(
                "ExampleResourceGroup",
                "ExampleWorkspace",
                "ExampleSqlPool",
                new MetadataSyncConfigInner().withEnabled(true),
                Context.NONE);
    }
}
