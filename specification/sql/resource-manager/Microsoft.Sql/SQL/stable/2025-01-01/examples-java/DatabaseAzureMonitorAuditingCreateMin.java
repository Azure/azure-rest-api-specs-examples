
import com.azure.resourcemanager.sql.fluent.models.DatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyName;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/**
 * Samples for DatabaseBlobAuditingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseAzureMonitorAuditingCreateMin.json
     */
    /**
     * Sample code: Create or update a database's azure monitor auditing policy with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateADatabaseSAzureMonitorAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseBlobAuditingPolicies().createOrUpdateWithResponse("blobauditingtest-4799",
            "blobauditingtest-6440", "testdb", BlobAuditingPolicyName.DEFAULT, new DatabaseBlobAuditingPolicyInner()
                .withIsAzureMonitorTargetEnabled(true).withState(BlobAuditingPolicyState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
