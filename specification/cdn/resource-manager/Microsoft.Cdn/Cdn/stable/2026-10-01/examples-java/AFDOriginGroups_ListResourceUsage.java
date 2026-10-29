
/**
 * Samples for AfdOriginGroups ListResourceUsage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDOriginGroups_ListResourceUsage.json
     */
    /**
     * Sample code: AFDOriginGroups_ListResourceUsage.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDOriginGroupsListResourceUsage(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdOriginGroups().listResourceUsage("RG", "profile1", "origingroup1",
            com.azure.core.util.Context.NONE);
    }
}
