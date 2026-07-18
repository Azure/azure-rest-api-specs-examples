
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyName;

/**
 * Samples for ServerBlobAuditingPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ServerBlobAuditingGet.json
     */
    /**
     * Sample code: Get a server's blob auditing policy.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAServerSBlobAuditingPolicy(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getServerBlobAuditingPolicies().getWithResponse("blobauditingtest-4799",
            "blobauditingtest-6440", BlobAuditingPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
