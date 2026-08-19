
/**
 * Samples for RoleMappings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-03-preview/RoleMappings_List_MinimumSet_Gen.json
     */
    /**
     * Sample code: RoleMappings_List_MinimumSet - List role mappings with single role.
     * 
     * @param manager Entry point to CommvaultContentStoreManager.
     */
    public static void roleMappingsListMinimumSetListRoleMappingsWithSingleRole(
        com.azure.resourcemanager.commvaultcontentstore.CommvaultContentStoreManager manager) {
        manager.roleMappings().list("rgcommvault", "myCloudAccount", com.azure.core.util.Context.NONE);
    }
}
