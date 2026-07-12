
/**
 * Samples for Extensions Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-30/DeleteExtension.json
     */
    /**
     * Sample code: Delete Arc Extension.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteArcExtension(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.extensions().delete("test-rg", "myCluster", "default", "MicrosoftMonitoringAgent",
            com.azure.core.util.Context.NONE);
    }
}
