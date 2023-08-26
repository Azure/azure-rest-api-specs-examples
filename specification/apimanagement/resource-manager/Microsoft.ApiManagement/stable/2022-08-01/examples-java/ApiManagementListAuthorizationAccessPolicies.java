/** Samples for AuthorizationAccessPolicy ListByAuthorization. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListAuthorizationAccessPolicies.json
     */
    /**
     * Sample code: ApiManagementListAuthorizationAccessPolicies.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListAuthorizationAccessPolicies(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationAccessPolicies()
            .listByAuthorization(
                "rg1", "apimService1", "aadwithauthcode", "authz1", null, null, null, com.azure.core.util.Context.NONE);
    }
}
