
/**
 * Samples for TagOperationLink Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetTagOperationLink.json
     */
    /**
     * Sample code: ApiManagementGetTagOperationLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetTagOperationLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tagOperationLinks().getWithResponse("rg1", "apimService1", "tag1", "link1",
            com.azure.core.util.Context.NONE);
    }
}
