
/**
 * Samples for TagApiLink Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteTagApiLink.json
     */
    /**
     * Sample code: ApiManagementDeleteTagApiLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteTagApiLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tagApiLinks().deleteWithResponse("rg1", "apimService1", "tag1", "link1",
            com.azure.core.util.Context.NONE);
    }
}
