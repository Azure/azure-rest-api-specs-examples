
import com.azure.resourcemanager.sql.fluent.models.ServerBlobAuditingPolicyInner;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyName;
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyState;

/**
 * Samples for ServerBlobAuditingPolicies CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerBlobAuditingCreateMin.json
     */
    /**
     * Sample code: Update a server's blob auditing policy with minimal parameters.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        updateAServerSBlobAuditingPolicyWithMinimalParameters(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerBlobAuditingPolicies().createOrUpdate("blobauditingtest-4799",
            "blobauditingtest-6440", BlobAuditingPolicyName.DEFAULT,
            new ServerBlobAuditingPolicyInner().withState(BlobAuditingPolicyState.ENABLED)
                .withStorageEndpoint("https://mystorage.blob.core.windows.net")
                .withStorageAccountAccessKey("fakeTokenPlaceholder"),
            com.azure.core.util.Context.NONE);
    }
}
