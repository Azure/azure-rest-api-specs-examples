
/**
 * Samples for ArcSettings Create.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/PutArcSetting.json
     */
    /**
     * Sample code: Create ArcSetting.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void createArcSetting(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.arcSettings().define("default").withExistingCluster("test-rg", "myCluster").create();
    }
}
