/** Samples for SqlPoolBlobAuditingPolicies ListBySqlPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/SqlPoolAuditingSettingsList.json
     */
    /**
     * Sample code: List audit settings of a database.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void listAuditSettingsOfADatabase(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .sqlPoolBlobAuditingPolicies()
            .listBySqlPool(
                "blobauditingtest-6852", "blobauditingtest-2080", "testdb", com.azure.core.util.Context.NONE);
    }
}
