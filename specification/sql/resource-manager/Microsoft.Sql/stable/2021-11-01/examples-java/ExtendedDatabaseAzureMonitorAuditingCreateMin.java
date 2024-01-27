
import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.ExtendedDatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/** Samples for ExtendedDatabaseBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ExtendedDatabaseAzureMonitorAuditingCreateMin.json
     */
    /**
     * Sample code: Create or update an extended database's azure monitor auditing policy with minimal parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateAnExtendedDatabaseSAzureMonitorAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getExtendedDatabaseBlobAuditingPolicies()
            .createOrUpdateWithResponse("blobauditingtest-4799", "blobauditingtest-6440", "testdb",
                new ExtendedDatabaseBlobAuditingPolicyInner().withIsAzureMonitorTargetEnabled(true)
                    .withState(BlobAuditingPolicyState.ENABLED),
                Context.NONE);
    }
}
