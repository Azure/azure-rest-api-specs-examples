/** Samples for AuthorizationProvider Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetAuthorizationProvider.json
     */
    /**
     * Sample code: ApiManagementGetAuthorizationProvider.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetAuthorizationProvider(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationProviders()
            .getWithResponse("rg1", "apimService1", "aadwithauthcode", com.azure.core.util.Context.NONE);
    }
}
