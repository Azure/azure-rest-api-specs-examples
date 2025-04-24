
/**
 * Samples for TagProductLink Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetTagProductLink.json
     */
    /**
     * Sample code: ApiManagementGetTagProductLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetTagProductLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tagProductLinks().getWithResponse("rg1", "apimService1", "tag1", "link1",
            com.azure.core.util.Context.NONE);
    }
}
