
/**
 * Samples for ApiWiki GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadApiWiki.json
     */
    /**
     * Sample code: ApiManagementHeadApiWiki.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadApiWiki(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiWikis().getEntityTagWithResponse("rg1", "apimService1", "57d1f7558aa04f15146d9d8a",
            com.azure.core.util.Context.NONE);
    }
}
