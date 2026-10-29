
/**
 * Samples for Origins ListByEndpoint.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Origins_ListByEndpoint.json
     */
    /**
     * Sample code: Origins_ListByEndpoint.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void originsListByEndpoint(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getOrigins().listByEndpoint("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
