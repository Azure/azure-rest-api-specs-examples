import com.azure.core.util.Context;

/** Samples for ContentItem ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListContentTypeContentItems.json
     */
    /**
     * Sample code: ApiManagementListContentTypeContentItems.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListContentTypeContentItems(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.contentItems().listByService("rg1", "apimService1", "page", Context.NONE);
    }
}
