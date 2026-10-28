
/**
 * Samples for Origins Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/Origins_Delete.json
     */
    /**
     * Sample code: Origins_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void originsDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getOrigins().delete("RG", "profile1", "endpoint1", "origin1",
            com.azure.core.util.Context.NONE);
    }
}
