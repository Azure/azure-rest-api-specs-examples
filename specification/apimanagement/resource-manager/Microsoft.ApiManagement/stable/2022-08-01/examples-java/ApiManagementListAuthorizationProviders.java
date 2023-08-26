/** Samples for AuthorizationProvider ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListAuthorizationProviders.json
     */
    /**
     * Sample code: ApiManagementListAuthorizationProviders.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListAuthorizationProviders(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationProviders()
            .listByService("rg1", "apimService1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
