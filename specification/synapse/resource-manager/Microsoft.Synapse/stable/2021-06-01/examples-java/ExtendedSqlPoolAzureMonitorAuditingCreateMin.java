import com.azure.resourcemanager.synapse.models.BlobAuditingPolicyState;

/** Samples for ExtendedSqlPoolBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/stable/2021-06-01/examples/ExtendedSqlPoolAzureMonitorAuditingCreateMin.json
     */
    /**
     * Sample code: Create or update an extended SQL pool's azure monitor auditing policy with minimal parameters.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void createOrUpdateAnExtendedSQLPoolSAzureMonitorAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .extendedSqlPoolBlobAuditingPolicies()
            .define()
            .withExistingSqlPool("blobauditingtest-4799", "blobauditingtest-6440", "testdb")
            .withState(BlobAuditingPolicyState.ENABLED)
            .withIsAzureMonitorTargetEnabled(true)
            .create();
    }
}
