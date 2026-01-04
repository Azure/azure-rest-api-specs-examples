
/**
 * Samples for WebApps ListSlots.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListWebAppSlots.json
     */
    /**
     * Sample code: List Web App Slots.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listWebAppSlots(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listSlots("testrg123", "sitef6141",
            com.azure.core.util.Context.NONE);
    }
}
