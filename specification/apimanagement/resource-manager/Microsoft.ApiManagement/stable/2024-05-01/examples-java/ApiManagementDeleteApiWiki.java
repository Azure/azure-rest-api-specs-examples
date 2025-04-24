
/**
 * Samples for ApiWiki Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteApiWiki.json
     */
    /**
     * Sample code: ApiManagementDeleteApiWiki.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteApiWiki(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiWikis().deleteWithResponse("rg1", "apimService1", "57d1f7558aa04f15146d9d8a", "*",
            com.azure.core.util.Context.NONE);
    }
}
