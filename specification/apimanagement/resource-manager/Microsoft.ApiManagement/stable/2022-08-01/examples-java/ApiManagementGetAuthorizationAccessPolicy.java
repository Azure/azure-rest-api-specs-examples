/** Samples for AuthorizationAccessPolicy Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetAuthorizationAccessPolicy.json
     */
    /**
     * Sample code: ApiManagementGetAuthorizationAccessPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetAuthorizationAccessPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationAccessPolicies()
            .getWithResponse(
                "rg1",
                "apimService1",
                "aadwithauthcode",
                "authz1",
                "fe0bed83-631f-4149-bd0b-0464b1bc7cab",
                com.azure.core.util.Context.NONE);
    }
}
