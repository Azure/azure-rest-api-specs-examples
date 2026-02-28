
/**
 * Samples for WebApps ListSiteBackupsSlot.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/ListSlotBackups.json
     */
    /**
     * Sample code: List backups.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listBackups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listSiteBackupsSlot("testrg123", "tests346", "staging",
            com.azure.core.util.Context.NONE);
    }
}
