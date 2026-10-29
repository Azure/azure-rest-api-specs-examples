
/**
 * Samples for Endpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Endpoints_Delete.json
     */
    /**
     * Sample code: Endpoints_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void endpointsDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getEndpoints().delete("RG", "profile1", "endpoint1", com.azure.core.util.Context.NONE);
    }
}
