
/**
 * Samples for AfdOriginGroups Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDOriginGroups_Delete.json
     */
    /**
     * Sample code: AFDOriginGroups_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDOriginGroupsDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdOriginGroups().delete("RG", "profile1", "origingroup1",
            com.azure.core.util.Context.NONE);
    }
}
