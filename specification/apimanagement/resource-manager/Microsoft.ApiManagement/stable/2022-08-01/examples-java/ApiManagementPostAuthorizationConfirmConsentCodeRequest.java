import com.azure.resourcemanager.apimanagement.models.AuthorizationConfirmConsentCodeRequestContract;

/** Samples for Authorization ConfirmConsentCode. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementPostAuthorizationConfirmConsentCodeRequest.json
     */
    /**
     * Sample code: ApiManagementPostAuthorizationConfirmConsentCodeRequest.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementPostAuthorizationConfirmConsentCodeRequest(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .authorizations()
            .confirmConsentCodeWithResponse(
                "rg1",
                "apimService1",
                "aadwithauthcode",
                "authz1",
                new AuthorizationConfirmConsentCodeRequestContract().withConsentCode("fakeTokenPlaceholder"),
                com.azure.core.util.Context.NONE);
    }
}
