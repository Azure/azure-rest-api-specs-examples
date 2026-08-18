
/**
 * Samples for RoleMappings List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-03-preview/RoleMappings_List_MaximumSet_Gen.json
     */
    /**
     * Sample code: RoleMappings_List.
     * 
     * @param manager Entry point to CommvaultContentStoreManager.
     */
    public static void
        roleMappingsList(com.azure.resourcemanager.commvaultcontentstore.CommvaultContentStoreManager manager) {
        manager.roleMappings().list("rgcommvault", "myCloudAccount", com.azure.core.util.Context.NONE);
    }
}
