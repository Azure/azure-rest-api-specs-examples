/** Samples for AuthorizationAccessPolicy CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateAuthorizationAccessPolicy.json
     */
    /**
     * Sample code: ApiManagementCreateAuthorizationAccessPolicy.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateAuthorizationAccessPolicy(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationAccessPolicies()
            .define("fe0bed83-631f-4149-bd0b-0464b1bc7cab")
            .withExistingAuthorization("rg1", "apimService1", "aadwithauthcode", "authz1")
            .withTenantId("13932a0d-5c63-4d37-901d-1df9c97722ff")
            .withObjectId("fe0bed83-631f-4149-bd0b-0464b1bc7cab")
            .create();
    }
}
