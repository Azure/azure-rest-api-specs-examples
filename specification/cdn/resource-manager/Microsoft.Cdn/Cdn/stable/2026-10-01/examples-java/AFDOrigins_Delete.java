
/**
 * Samples for AfdOrigins Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01/AFDOrigins_Delete.json
     */
    /**
     * Sample code: AFDOrigins_Delete.
     * 
     * @param manager Entry point to CdnManager.
     */
    public static void aFDOriginsDelete(com.azure.resourcemanager.cdn.CdnManager manager) {
        manager.serviceClient().getAfdOrigins().delete("RG", "profile1", "origingroup1", "origin1",
            com.azure.core.util.Context.NONE);
    }
}
