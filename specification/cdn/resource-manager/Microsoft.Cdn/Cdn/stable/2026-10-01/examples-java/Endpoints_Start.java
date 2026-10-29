
/**
 * Samples for Endpoints Start.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Endpoints_Start.json
     */
    /**
     * Sample code: Endpoints_Start.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void endpointsStart(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getEndpoints().start("RG", "profile1", "endpoint1", com.azure.core.util.Context.NONE);
    }
}
