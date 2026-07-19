
import com.azure.resourcemanager.sql.fluent.models.ExtendedDatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyName;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/**
 * Samples for ExtendedDatabaseBlobAuditingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ExtendedDatabaseAzureMonitorAuditingCreateMin.json
     */
    /**
     * Sample code: Create or update an extended database's azure monitor auditing policy with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateAnExtendedDatabaseSAzureMonitorAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getExtendedDatabaseBlobAuditingPolicies()
            .createOrUpdateWithResponse("blobauditingtest-4799", "blobauditingtest-6440", "testdb",
                BlobAuditingPolicyName.DEFAULT, new ExtendedDatabaseBlobAuditingPolicyInner()
                    .withIsAzureMonitorTargetEnabled(true).withState(BlobAuditingPolicyState.ENABLED),
                com.azure.core.util.Context.NONE);
    }
}
