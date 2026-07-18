
import com.azure.resourcemanager.sql.models.BlobAuditingPolicyName;

/**
 * Samples for ExtendedServerBlobAuditingPolicies Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ExtendedServerBlobAuditingGet.json
     */
    /**
     * Sample code: Get a server's blob extended auditing policy.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getAServerSBlobExtendedAuditingPolicy(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getExtendedServerBlobAuditingPolicies().getWithResponse("blobauditingtest-4799",
            "blobauditingtest-6440", BlobAuditingPolicyName.DEFAULT, com.azure.core.util.Context.NONE);
    }
}
