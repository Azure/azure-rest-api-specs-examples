
import com.azure.resourcemanager.sql.fluent.models.ExtendedServerBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyName;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/**
 * Samples for ExtendedServerBlobAuditingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ExtendedServerBlobAuditingCreateMin.json
     */
    /**
     * Sample code: Update a server's extended blob auditing policy with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void updateAServerSExtendedBlobAuditingPolicyWithMinimalParameters(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getExtendedServerBlobAuditingPolicies().createOrUpdate("blobauditingtest-4799",
            "blobauditingtest-6440", BlobAuditingPolicyName.DEFAULT,
            new ExtendedServerBlobAuditingPolicyInner().withState(BlobAuditingPolicyState.ENABLED)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
