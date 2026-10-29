
/**
 * Samples for Routes ListByEndpoint.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Routes_ListByEndpoint.json
     */
    /**
     * Sample code: Routes_ListByEndpoint.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void routesListByEndpoint(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRoutes().listByEndpoint("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
