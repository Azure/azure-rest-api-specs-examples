import com.azure.core.util.Context;

/** Samples for WebApps ListSiteBackups. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2021-03-01/examples/ListSlotBackups.json
     */
    /**
     * Sample code: List backups.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listBackups(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().listSiteBackups("testrg123", "tests346", Context.NONE);
    }
}
