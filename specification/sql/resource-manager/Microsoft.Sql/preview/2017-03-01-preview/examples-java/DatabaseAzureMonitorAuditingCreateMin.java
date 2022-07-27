import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/** Samples for DatabaseBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/DatabaseAzureMonitorAuditingCreateMin.json
     */
    /**
     * Sample code: Create or update a database's azure monitor auditing policy with minimal parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateADatabaseSAzureMonitorAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getDatabaseBlobAuditingPolicies()
            .createOrUpdateWithResponse(
                "blobauditingtest-4799",
                "blobauditingtest-6440",
                "testdb",
                new DatabaseBlobAuditingPolicyInner()
                    .withState(BlobAuditingPolicyState.ENABLED)
                    .withIsAzureMonitorTargetEnabled(true),
                Context.NONE);
    }
}
