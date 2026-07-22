
/**
 * Samples for AccessPolicyAssignment Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-06-01-preview/RedisEnterpriseAccessPolicyAssignmentGetWithFailedProvisioningState.json
     */
    /**
     * Sample code: RedisEnterpriseAccessPolicyAssignmentGetWithFailedProvisioningState.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseAccessPolicyAssignmentGetWithFailedProvisioningState(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.accessPolicyAssignments().getWithResponse("rg1", "cache1", "default", "accessPolicyAssignmentName1",
            com.azure.core.util.Context.NONE);
    }
}
