import com.azure.resourcemanager.apimanagement.models.IdentityProviderType;

/** Samples for IdentityProvider Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementGetIdentityProvider.json
     */
    /**
     * Sample code: ApiManagementGetIdentityProvider.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementGetIdentityProvider(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .identityProviders()
            .getWithResponse("rg1", "apimService1", IdentityProviderType.AAD_B2C, com.azure.core.util.Context.NONE);
    }
}
