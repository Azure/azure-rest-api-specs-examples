
/**
 * Samples for AccessPolicyAssignment Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RedisEnterpriseAccessPolicyAssignmentGet.json
     */
    /**
     * Sample code: RedisEnterpriseAccessPolicyAssignmentGet.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseAccessPolicyAssignmentGet(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.accessPolicyAssignments().getWithResponse("rg1", "cache1", "default", "accessPolicyAssignmentName1",
            com.azure.core.util.Context.NONE);
    }
}
