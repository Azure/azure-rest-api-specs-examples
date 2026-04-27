
/**
 * Samples for AppLinks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/AppLinks_ListByResourceGroup.json
     */
    /**
     * Sample code: AppLinks_ListByResourceGroup.
     * 
     * @param manager Entry point to AppnetworkManager.
     */
    public static void appLinksListByResourceGroup(com.azure.resourcemanager.appnetwork.AppnetworkManager manager) {
        manager.appLinks().listByResourceGroup("test_rg", com.azure.core.util.Context.NONE);
    }
}
