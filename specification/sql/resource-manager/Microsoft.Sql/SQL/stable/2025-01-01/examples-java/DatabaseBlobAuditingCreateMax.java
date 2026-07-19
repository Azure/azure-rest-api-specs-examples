
import com.azure.resourcemanager.sql.fluent.models.DatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyName;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;
import java.util.Arrays;
import java.util.UUID;

/**
 * Samples for DatabaseBlobAuditingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/DatabaseBlobAuditingCreateMax.json
     */
    /**
     * Sample code: Create or update a database's blob auditing policy with all parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void createOrUpdateADatabaseSBlobAuditingPolicyWithAllParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getDatabaseBlobAuditingPolicies().createOrUpdateWithResponse("blobauditingtest-4799",
            "blobauditingtest-6440", "testdb", BlobAuditingPolicyName.DEFAULT,
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
