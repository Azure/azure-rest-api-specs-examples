
/**
 * Samples for TagApiLink Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementGetTagApiLink.json
     */
    /**
     * Sample code: ApiManagementGetTagApiLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementGetTagApiLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tagApiLinks().getWithResponse("rg1", "apimService1", "tag1", "link1", com.azure.core.util.Context.NONE);
    }
}
