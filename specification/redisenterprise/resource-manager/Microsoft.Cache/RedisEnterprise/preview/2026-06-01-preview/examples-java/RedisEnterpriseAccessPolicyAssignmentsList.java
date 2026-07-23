
/**
 * Samples for AccessPolicyAssignment List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RedisEnterpriseAccessPolicyAssignmentsList.json
     */
    /**
     * Sample code: RedisEnterpriseAccessPolicyAssignmentList.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseAccessPolicyAssignmentList(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.accessPolicyAssignments().list("rg1", "cache1", "default", com.azure.core.util.Context.NONE);
    }
}
