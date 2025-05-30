
/**
 * Samples for ContentItem GetEntityTag.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadContentTypeContentItem.json
     */
    /**
     * Sample code: ApiManagementHeadContentTypeContentItem.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadContentTypeContentItem(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.contentItems().getEntityTagWithResponse("rg1", "apimService1", "page",
            "4e3cf6a5-574a-ba08-1f23-2e7a38faa6d8", com.azure.core.util.Context.NONE);
    }
}
