
/**
 * Samples for AppLinks Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-01-preview/AppLinks_Delete.json
     */
    /**
     * Sample code: AppLinks_Delete.
     * 
     * @param manager Entry point to AppnetworkManager.
     */
    public static void appLinksDelete(com.azure.resourcemanager.appnetwork.AppnetworkManager manager) {
        manager.appLinks().delete("test_rg", "applink-test-01", com.azure.core.util.Context.NONE);
    }
}
