
/**
 * Samples for Endpoints Stop.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Endpoints_Stop.json
     */
    /**
     * Sample code: Endpoints_Stop.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void endpointsStop(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getEndpoints().stop("RG", "profile1", "endpoint1", com.azure.core.util.Context.NONE);
    }
}
