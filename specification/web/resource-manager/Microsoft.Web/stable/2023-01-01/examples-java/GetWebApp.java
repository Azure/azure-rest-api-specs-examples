/** Samples for WebApps GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/GetWebApp.json
     */
    /**
     * Sample code: Get Web App.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getWebApp(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .webApps()
            .manager()
            .serviceClient()
            .getWebApps()
            .getByResourceGroupWithResponse("testrg123", "sitef6141", com.azure.core.util.Context.NONE);
    }
}
