
import com.azure.resourcemanager.redisenterprise.models.AccessPolicyAssignmentPropertiesUser;

/**
 * Samples for AccessPolicyAssignment CreateUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/redisenterprise/resource-manager/Microsoft.Cache/preview/2024-09-01-preview/examples/
     * RedisEnterpriseAccessPolicyAssignmentCreateUpdate.json
     */
    /**
     * Sample code: RedisEnterpriseAccessPolicyAssignmentCreateUpdate.
     * 
     * @param manager Entry point to RedisEnterpriseManager.
     */
    public static void redisEnterpriseAccessPolicyAssignmentCreateUpdate(
        com.azure.resourcemanager.redisenterprise.RedisEnterpriseManager manager) {
        manager.accessPolicyAssignments().define("defaultTestEntraApp1")
            .withExistingDatabase("rg1", "cache1", "default").withAccessPolicyName("default")
            .withUser(new AccessPolicyAssignmentPropertiesUser().withObjectId("6497c918-11ad-41e7-1b0f-7c518a87d0b0"))
            .create();
    }
}
