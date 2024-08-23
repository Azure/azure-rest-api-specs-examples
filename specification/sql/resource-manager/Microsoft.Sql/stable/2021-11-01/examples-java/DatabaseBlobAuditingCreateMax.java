
import com.azure.resourcemanager.sql.fluent.models.DatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;
import java.util.Arrays;
import java.util.UUID;

/**
 * Samples for DatabaseBlobAuditingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/DatabaseBlobAuditingCreateMax.json
     */
    /**
     * Sample code: Create or update a database's blob auditing policy with all parameters.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateADatabaseSBlobAuditingPolicyWithAllParameters(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getDatabaseBlobAuditingPolicies().createOrUpdateWithResponse(
            "blobauditingtest-4799", "blobauditingtest-6440", "testdb",
            new DatabaseBlobAuditingPolicyInner().withRetentionDays(6)
                .withAuditActionsAndGroups(Arrays.asList("DATABASE_LOGOUT_GROUP", "DATABASE_ROLE_MEMBER_CHANGE_GROUP",
                    "UPDATE on database::TestDatabaseName by public"))
                .withIsStorageSecondaryKeyInUse(false).withIsAzureMonitorTargetEnabled(true).withQueueDelayMs(4000)
                .withState(BlobAuditingPolicyState.ENABLED)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder").withStorageAccountSubscriptionId(
                    UUID.fromString("00000000-1234-0000-5678-000000000000")),
            com.azure.core.util.Context.NONE);
    }
}
