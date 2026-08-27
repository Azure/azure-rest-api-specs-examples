
/**
 * Samples for HorizonDbPrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/PrivateLinkResources_List.json
     */
    /**
     * Sample code: Gets private link resources for HorizonDB.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        getsPrivateLinkResourcesForHorizonDB(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.horizonDbPrivateLinkResources().list("exampleresourcegroup", "examplecluster",
            com.azure.core.util.Context.NONE);
    }
}
