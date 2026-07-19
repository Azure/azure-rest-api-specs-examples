
import com.azure.resourcemanager.sql.fluent.models.ExtendedServerBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyName;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;
import java.util.Arrays;
import java.util.UUID;

/**
 * Samples for ExtendedServerBlobAuditingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ExtendedServerBlobAuditingCreateMax.json
     */
    /**
     * Sample code: Update a server's extended blob auditing policy with all parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAServerSExtendedBlobAuditingPolicyWithAllParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getExtendedServerBlobAuditingPolicies().createOrUpdate("blobauditingtest-4799",
            "blobauditingtest-6440", BlobAuditingPolicyName.DEFAULT,
            new ExtendedServerBlobAuditingPolicyInner().withPredicateExpression("object_name = 'SensitiveData'")
                .withRetentionDays(6)
                .withAuditActionsAndGroups(Arrays.asList("SUCCESSFUL_DATABASE_AUTHENTICATION_GROUP",
                    "FAILED_DATABASE_AUTHENTICATION_GROUP", "BATCH_COMPLETED_GROUP"))
                .withIsStorageSecondaryKeyInUse(false).withIsAzureMonitorTargetEnabled(true).withQueueDelayMs(4000)
                .withState(BlobAuditingPolicyState.ENABLED)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder").withStorageAccountSubscriptionId(
                    UUID.fromString("00000000-1234-0000-5678-000000000000")),
            com.azure.core.util.Context.NONE);
    }
}
