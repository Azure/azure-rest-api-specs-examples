
/**
 * Samples for GroupUser CheckEntityExists.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementHeadGroupUser.json
     */
    /**
     * Sample code: ApiManagementHeadGroupUser.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementHeadGroupUser(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.groupUsers().checkEntityExistsWithResponse("rg1", "apimService1", "59306a29e4bbd510dc24e5f9",
            "5931a75ae4bbd512a88c680b", com.azure.core.util.Context.NONE);
    }
}
