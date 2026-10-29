
/**
 * Samples for AfdEndpoints ListByProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDEndpoints_ListByProfile.json
     */
    /**
     * Sample code: AFDEndpoints_ListByProfile.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDEndpointsListByProfile(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdEndpoints().listByProfile("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
