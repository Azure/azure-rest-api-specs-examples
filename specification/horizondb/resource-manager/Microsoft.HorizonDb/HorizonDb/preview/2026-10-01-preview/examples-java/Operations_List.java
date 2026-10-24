
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/Operations_List.json
     */
    /**
     * Sample code: List operations for Microsoft.HorizonDb.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listOperationsForMicrosoftHorizonDb(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
