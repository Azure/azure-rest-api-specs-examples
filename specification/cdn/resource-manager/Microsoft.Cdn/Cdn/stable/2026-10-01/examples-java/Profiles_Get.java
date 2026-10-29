
/**
 * Samples for Profiles GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Profiles_Get.json
     */
    /**
     * Sample code: Profiles_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void profilesGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getProfiles().getByResourceGroupWithResponse("RG", "profile1",
            com.azure.core.util.Context.NONE);
    }
}
