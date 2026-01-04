
/**
 * Samples for WebApps ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListWebAppsByResourceGroup
     * .json
     */
    /**
     * Sample code: List Web Apps by Resource group.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listWebAppsByResourceGroup(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listByResourceGroup("testrg123", null,
            com.azure.core.util.Context.NONE);
    }
}
