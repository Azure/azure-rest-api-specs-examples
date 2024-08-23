
import com.azure.resourcemanager.sql.fluent.models.DatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/**
 * Samples for DatabaseBlobAuditingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseAzureMonitorAuditingCreateMin
     * .json
     */
    /**
     * Sample code: Create or update a database's azure monitor auditing policy with minimal parameters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateADatabaseSAzureMonitorAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseBlobAuditingPolicies().createOrUpdateWithResponse(
            "blobauditingtest-4799", "blobauditingtest-6440", "testdb", new DatabaseBlobAuditingPolicyInner()
                .withIsAzureMonitorTargetEnabled(true).withState(BlobAuditingPolicyState.ENABLED),
            com.azure.core.util.Context.NONE);
    }
}
