
/**
 * Samples for MicrosoftEntraAdministrators Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/MicrosoftEntraAdministrators_Delete.json
     */
    /**
     * Sample code: Delete a Microsoft Entra administrator from a HorizonDB cluster.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void deleteAMicrosoftEntraAdministratorFromAHorizonDBCluster(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.microsoftEntraAdministrators().delete("example-resource-group", "example-cluster",
            "11111111-1111-1111-1111-111111111111", com.azure.core.util.Context.NONE);
    }
}
