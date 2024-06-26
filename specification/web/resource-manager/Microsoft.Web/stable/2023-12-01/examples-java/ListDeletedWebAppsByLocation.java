
/**
 * Samples for DeletedWebApps ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-12-01/examples/ListDeletedWebAppsByLocation.json
     */
    /**
     * Sample code: List Deleted Web App by Location.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listDeletedWebAppByLocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDeletedWebApps().listByLocation("West US 2",
            com.azure.core.util.Context.NONE);
    }
}
