
/**
 * Samples for PrivateLinkResources List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/PrivateLinkResources_List.json
     */
    /**
     * Sample code: List HorizonDB private link resources.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listHorizonDBPrivateLinkResources(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.privateLinkResources().list("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
