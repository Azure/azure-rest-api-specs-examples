
/**
 * Samples for TagProductLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteTagProductLink.json
     */
    /**
     * Sample code: ApiManagementDeleteTagProductLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteTagProductLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tagProductLinks().deleteWithResponse("rg1", "apimService1", "tag1", "link1",
            com.azure.core.util.Context.NONE);
    }
}
