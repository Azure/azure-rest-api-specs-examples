
/**
 * Samples for ArcSettings Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-02-01/DeleteArcSetting.json
     */
    /**
     * Sample code: Delete ArcSetting.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void deleteArcSetting(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.arcSettings().delete("test-rg", "myCluster", "default", com.azure.core.util.Context.NONE);
    }
}
