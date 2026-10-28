
/**
 * Samples for Endpoints ListByProfile.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Endpoints_ListByProfile.json
     */
    /**
     * Sample code: Endpoints_ListByProfile.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void endpointsListByProfile(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getEndpoints().listByProfile("RG", "profile1", com.azure.core.util.Context.NONE);
    }
}
