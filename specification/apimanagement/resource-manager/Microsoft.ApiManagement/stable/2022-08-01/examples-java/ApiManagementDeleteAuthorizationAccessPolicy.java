
/**
 * Samples for AuthorizationAccessPolicy Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/
     * ApiManagementDeleteAuthorizationAccessPolicy.json
     */
    /**
     * Sample code: ApiManagementDeleteAuthorizationAccessPolicy.
     * 
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteAuthorizationAccessPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.authorizationAccessPolicies().deleteWithResponse("rg1", "apimService1", "aadwithauthcode", "authz1",
            "fe0bed83-631f-4149-bd0b-0464b1bc7cab", "*", com.azure.core.util.Context.NONE);
    }
}
