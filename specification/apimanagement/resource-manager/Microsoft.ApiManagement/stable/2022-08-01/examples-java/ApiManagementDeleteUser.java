/** Samples for User Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteUser.json
     */
    /**
     * Sample code: ApiManagementDeleteUser.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteUser(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .users()
            .deleteWithResponse(
                "rg1",
                "apimService1",
                "5931a75ae4bbd512288c680b",
                "*",
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
