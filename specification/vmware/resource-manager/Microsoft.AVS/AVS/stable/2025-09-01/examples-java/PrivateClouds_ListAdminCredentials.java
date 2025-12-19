
/**
 * Samples for PrivateClouds ListAdminCredentials.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01/PrivateClouds_ListAdminCredentials.json
     */
    /**
     * Sample code: PrivateClouds_ListAdminCredentials.
     * 
     * @param manager Entry point to AvsManager.
     */
    public static void privateCloudsListAdminCredentials(com.azure.resourcemanager.avs.AvsManager manager) {
        manager.privateClouds().listAdminCredentialsWithResponse("group1", "cloud1", com.azure.core.util.Context.NONE);
    }
}
