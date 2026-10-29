
/**
 * Samples for Endpoints ListResourceUsage.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Endpoints_ListResourceUsage.json
     */
    /**
     * Sample code: Endpoints_ListResourceUsage.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void endpointsListResourceUsage(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getEndpoints().listResourceUsage("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
