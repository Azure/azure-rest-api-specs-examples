
/**
 * Samples for Extensions ListByArcSetting.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-01-preview/ListExtensionsByArcSetting.json
     */
    /**
     * Sample code: List Extensions under ArcSetting resource.
     * 
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void
        listExtensionsUnderArcSettingResource(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager.extensions().listByArcSetting("test-rg", "myCluster", "default", com.azure.core.util.Context.NONE);
    }
}
