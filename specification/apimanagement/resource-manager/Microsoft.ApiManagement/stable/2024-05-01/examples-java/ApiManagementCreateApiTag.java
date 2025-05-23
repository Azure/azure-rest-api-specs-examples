
/**
 * Samples for Tag AssignToApi.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateApiTag.json
     */
    /**
     * Sample code: ApiManagementCreateApiTag.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateApiTag(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tags().assignToApiWithResponse("rg1", "apimService1", "5931a75ae4bbd512a88c680b", "tagId1",
            com.azure.core.util.Context.NONE);
    }
}
