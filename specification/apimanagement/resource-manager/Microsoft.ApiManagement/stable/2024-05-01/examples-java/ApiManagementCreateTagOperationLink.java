
/**
 * Samples for TagOperationLink CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementCreateTagOperationLink.json
     */
    /**
     * Sample code: ApiManagementCreateTagOperationLink.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementCreateTagOperationLink(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.tagOperationLinks().define("link1").withExistingTag("rg1", "apimService1", "tag1").withOperationId(
            "/subscriptions/00000000-0000-0000-0000-000000000000/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/apis/echo-api/operations/op1")
            .create();
    }
}
