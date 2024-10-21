
/**
 * Samples for AccessPolicyAssignment Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/
     * RedisEnterpriseAccessPolicyAssignmentDelete.json
     */
    /**
     * Sample code: RedisEnterpriseAccessPolicyAssignmentDelete.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseAccessPolicyAssignmentDelete(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.accessPolicyAssignments().delete("rg1", "cache1", "default", "defaultTestEntraApp1",
            com.azure.core.util.Context.NONE);
    }
}
