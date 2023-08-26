/** Samples for User GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadUser.json
     */
    /**
     * Sample code: ApiManagementHeadUser.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadUser(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .users()
            .getEntityTagWithResponse(
                "rg1", "apimService1", "5931a75ae4bbd512a88c680b", com.azure.core.util.Context.NONE);
    }
}
