
/**
 * Samples for AccessPolicyAssignment List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/
     * RedisEnterpriseAccessPolicyAssignmentsList.json
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
