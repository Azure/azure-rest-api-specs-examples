
/**
 * Samples for TagOperationLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteTagOperationLink.json
     */
    /**
     * Sample code: ApiManagementDeleteTagOperationLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteTagOperationLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tagOperationLinks().deleteWithResponse("rg1", "apimService1", "tag1", "link1",
            com.azure.core.util.Context.NONE);
    }
}
