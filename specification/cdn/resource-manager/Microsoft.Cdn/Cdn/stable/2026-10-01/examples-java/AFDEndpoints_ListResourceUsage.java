
/**
 * Samples for AfdEndpoints ListResourceUsage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDEndpoints_ListResourceUsage.json
     */
    /**
     * Sample code: AFDEndpoints_ListResourceUsage.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDEndpointsListResourceUsage(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdEndpoints().listResourceUsage("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
