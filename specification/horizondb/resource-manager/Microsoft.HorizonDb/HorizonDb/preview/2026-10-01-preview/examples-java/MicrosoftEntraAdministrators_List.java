
/**
 * Samples for MicrosoftEntraAdministrators List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MicrosoftEntraAdministrators_List.json
     */
    /**
     * Sample code: List Microsoft Entra administrators for a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listMicrosoftEntraAdministratorsForAHorizonDBCluster(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.microsoftEntraAdministrators().list("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
