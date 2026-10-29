
/**
 * Samples for AfdOriginGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDOriginGroups_Get.json
     */
    /**
     * Sample code: AFDOriginGroups_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDOriginGroupsGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdOriginGroups().getWithResponse("RG", "profile1", "origingroup1",
            com.azure.core.util.Context.NONE);
    }
}
