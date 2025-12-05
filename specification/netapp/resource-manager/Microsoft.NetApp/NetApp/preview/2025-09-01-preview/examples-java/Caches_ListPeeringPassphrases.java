
/**
 * Samples for Caches ListPeeringPassphrases.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Caches_ListPeeringPassphrases.json
     */
    /**
     * Sample code: Caches_ListPeeringPassphrases.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void cachesListPeeringPassphrases(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.caches().listPeeringPassphrasesWithResponse("myRG", "account1", "pool1", "cache-1",
            com.azure.core.util.Context.NONE);
    }
}
