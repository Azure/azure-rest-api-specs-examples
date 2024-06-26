import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ExtendedDatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/** Samples for ExtendedDatabaseBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/ExtendedDatabaseAzureMonitorAuditingCreateMin.json
     */
    /**
     * Sample code: Create or update an extended database's azure monitor auditing policy with minimal parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAnExtendedDatabaseSAzureMonitorAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .sqlServers()
            .manager()
            .serviceClient()
            .getExtendedDatabaseBlobAuditingPolicies()
            .createOrUpdateWithResponse(
                "blobauditingtest-4799",
                "blobauditingtest-6440",
                "testdb",
                new ExtendedDatabaseBlobAuditingPolicyInner()
                    .withState(BlobAuditingPolicyState.ENABLED)
                    .withIsAzureMonitorTargetEnabled(true),
                Context.NONE);
    }
}
