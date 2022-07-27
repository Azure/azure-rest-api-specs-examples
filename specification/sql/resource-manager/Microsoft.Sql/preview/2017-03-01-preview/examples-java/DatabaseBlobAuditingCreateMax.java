import com.azure.core.util.Context;
import com.azure.resourcemanager.sql.fluent.models.DatabaseBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;
import java.util.Arrays;
import java.util.UUID;

/** Samples for DatabaseBlobAuditingPolicies CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/preview/2017-03-01-preview/examples/DatabaseBlobAuditingCreateMax.json
     */
    /**
     * Sample code: Create or update a database's blob auditing policy with all parameters.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createOrUpdateADatabaseSBlobAuditingPolicyWithAllParameters(
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
                    .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                    .withStorageAccountAccessKey(
                        "sdlfkjabc+sdlfkjsdlkfsjdfLDKFTERLKFDFKLjsdfksjdflsdkfD2342309432849328476458/3RSD==")
                    .withRetentionDays(6)
                    .withAuditActionsAndGroups(
                        Arrays
                            .asList(
                                "DATABASE_LOGOUT_GROUP",
                                "DATABASE_ROLE_MEMBER_CHANGE_GROUP",
                                "UPDATE on database::TestDatabaseName by public"))
                    .withStorageAccountSubscriptionId(UUID.fromString("00000000-1234-0000-5678-000000000000"))
                    .withIsStorageSecondaryKeyInUse(false)
                    .withIsAzureMonitorTargetEnabled(true)
                    .withQueueDelayMs(4000),
                Context.NONE);
    }
}
