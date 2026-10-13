
/**
 * Samples for RoleMappings Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-30/RoleMappings_Get_MinimumSet_Gen.json
     */
    /**
     * Sample code: RoleMappings_Get_MinimumSet - Get role mappings with single role.
     * 
     * @param manager Entry point to CommvaultContentStoreManager.
     */
    public static void roleMappingsGetMinimumSetGetRoleMappingsWithSingleRole(
        com.azure.resourcemanager.commvaultcontentstore.CommvaultContentStoreManager manager) {
        manager.roleMappings().getWithResponse("rgcommvault", "myCloudAccount", com.azure.core.util.Context.NONE);
    }
}
