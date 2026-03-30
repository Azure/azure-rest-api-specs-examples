
/**
 * Samples for WebApps ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/ListWebAppsByResourceGroup.json
     */
    /**
     * Sample code: List Web Apps by Resource group.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void listWebAppsByResourceGroup(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().listByResourceGroup("testrg123", null, com.azure.core.util.Context.NONE);
    }
}
