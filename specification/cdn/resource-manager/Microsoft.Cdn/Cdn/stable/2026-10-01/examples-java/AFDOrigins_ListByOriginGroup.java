
/**
 * Samples for AfdOrigins ListByOriginGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDOrigins_ListByOriginGroup.json
     */
    /**
     * Sample code: AFDOrigins_ListByOriginGroup.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDOriginsListByOriginGroup(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdOrigins().listByOriginGroup("RG", "profile1", "origingroup1",
            com.azure.core.util.Context.NONE);
    }
}
