
/**
 * Samples for DeletedWebApps GetDeletedWebAppByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/AppService/stable/2025-03-01/examples/GetDeletedWebAppByLocation
     * .json
     */
    /**
     * Sample code: Get Deleted Web App by Location.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getDeletedWebAppByLocation(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getDeletedWebApps()
            .getDeletedWebAppByLocationWithResponse("West US 2", "9", com.azure.core.util.Context.NONE);
    }
}
