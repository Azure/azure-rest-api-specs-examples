import com.azure.resourcemanager.apimanagement.models.AuthorizationLoginRequestContract;

/** Samples for AuthorizationLoginLinks Post. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetAuthorizationLoginRequest.json
     */
    /**
     * Sample code: ApiManagementGetAuthorizationLoginRequest.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetAuthorizationLoginRequest(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizationLoginLinks()
            .postWithResponse(
                "rg1",
                "apimService1",
                "aadwithauthcode",
                "authz1",
                new AuthorizationLoginRequestContract().withPostLoginRedirectUrl("https://www.bing.com/"),
                com.azure.core.util.Context.NONE);
    }
}
