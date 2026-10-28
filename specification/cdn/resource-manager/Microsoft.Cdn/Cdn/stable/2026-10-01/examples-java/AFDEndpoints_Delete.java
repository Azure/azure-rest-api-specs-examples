
/**
 * Samples for AfdEndpoints Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDEndpoints_Delete.json
     */
    /**
     * Sample code: AFDEndpoints_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDEndpointsDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdEndpoints().delete("RG", "profile1", "endpoint1",
            com.azure.core.util.Context.NONE);
    }
}
