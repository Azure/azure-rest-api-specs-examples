
/**
 * Samples for ContentItem Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteContentTypeContentItem.json
     */
    /**
     * Sample code: ApiManagementDeleteContentTypeContentItem.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteContentTypeContentItem(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.contentItems().deleteWithResponse("rg1", "apimService1", "page", "4e3cf6a5-574a-ba08-1f23-2e7a38faa6d8",
            "*", com.azure.core.util.Context.NONE);
    }
}
