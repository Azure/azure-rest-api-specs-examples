
/**
 * Samples for TagOperationLink ListByProduct.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListTagOperationLinks.json
     */
    /**
     * Sample code: ApiManagementListTagOperationLinks.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListTagOperationLinks(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tagOperationLinks().listByProduct("rg1", "apimService1", "tag1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
