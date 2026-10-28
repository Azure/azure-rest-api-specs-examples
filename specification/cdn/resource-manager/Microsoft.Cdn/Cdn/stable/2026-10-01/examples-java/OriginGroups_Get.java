
/**
 * Samples for OriginGroups Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/OriginGroups_Get.json
     */
    /**
     * Sample code: OriginGroups_Get.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void originGroupsGet(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getOriginGroups().getWithResponse("RG", "profile1", "endpoint1", "originGroup1",
            com.azure.core.util.Context.NONE);
    }
}
