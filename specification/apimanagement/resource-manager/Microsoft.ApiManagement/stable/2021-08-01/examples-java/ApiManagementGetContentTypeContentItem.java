import com.azure.core.util.Context;

/** Samples for ContentItem Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementGetContentTypeContentItem.json
     */
    /**
     * Sample code: ApiManagementGetContentTypeContentItem.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetContentTypeContentItem(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .contentItems()
            .getWithResponse("rg1", "apimService1", "page", "4e3cf6a5-574a-ba08-1f23-2e7a38faa6d8", Context.NONE);
    }
}
