
/**
 * Samples for Routes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Routes_Delete.json
     */
    /**
     * Sample code: Routes_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void routesDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getRoutes().delete("RG", "profile1", "endpoint1", "route1",
            com.azure.core.util.Context.NONE);
    }
}
