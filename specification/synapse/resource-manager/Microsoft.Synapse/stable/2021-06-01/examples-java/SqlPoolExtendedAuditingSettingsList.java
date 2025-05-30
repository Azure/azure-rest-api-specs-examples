
/**
 * Samples for ExtendedSqlPoolBlobAuditingPolicies ListBySqlPool.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/
     * SqlPoolExtendedAuditingSettingsList.json
     */
    /**
     * Sample code: List extended auditing settings of a database.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void
        listExtendedAuditingSettingsOfADatabase(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.extendedSqlPoolBlobAuditingPolicies().listBySqlPool("blobauditingtest-6852", "blobauditingtest-2080",
            "testdb", com.azure.core.util.Context.NONE);
    }
}
