
/**
 * Samples for TagProductLink ListByProduct.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementListTagProductLinks.json
     */
    /**
     * Sample code: ApiManagementListTagProductLinks.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementListTagProductLinks(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tagProductLinks().listByProduct("rg1", "apimService1", "tag1", null, null, null,
            com.azure.core.util.Context.NONE);
    }
}
