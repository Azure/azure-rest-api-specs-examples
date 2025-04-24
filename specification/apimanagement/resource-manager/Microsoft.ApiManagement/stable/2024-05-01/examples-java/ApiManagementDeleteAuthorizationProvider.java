
/**
 * Samples for AuthorizationProvider Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2024-05-01/examples/
     * ApiManagementDeleteAuthorizationProvider.json
     */
    /**
     * Sample code: ApiManagementDeleteAuthorizationProvider.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void
        apiManagementDeleteAuthorizationProvider(com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.authorizationProviders().deleteWithResponse("rg1", "apimService1", "aadwithauthcode", "*",
            com.azure.core.util.Context.NONE);
    }
}
