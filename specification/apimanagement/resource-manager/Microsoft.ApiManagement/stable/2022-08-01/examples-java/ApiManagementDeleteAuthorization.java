/** Samples for Authorization Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteAuthorization.json
     */
    /**
     * Sample code: ApiManagementDeleteAuthorization.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteAuthorization(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizations()
            .deleteWithResponse(
                "rg1", "apimService1", "aadwithauthcode", "authz1", "*", com.azure.core.util.Context.NONE);
    }
}
