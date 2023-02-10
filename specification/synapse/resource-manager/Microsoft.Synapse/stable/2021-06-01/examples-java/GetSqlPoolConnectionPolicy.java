import com.azure.resourcemanager.synapse.models.ConnectionPolicyName;

/** Samples for SqlPoolConnectionPolicies Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/GetSqlPoolConnectionPolicy.json
     */
    /**
     * Sample code: Get a connection policy of a SQL Analytics pool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void getAConnectionPolicyOfASQLAnalyticsPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolConnectionPolicies()
            .getWithResponse(
                "blobauditingtest-6852",
                "blobauditingtest-2080",
                "testdb",
                ConnectionPolicyName.DEFAULT,
                com.azure.core.util.Context.NONE);
    }
}
