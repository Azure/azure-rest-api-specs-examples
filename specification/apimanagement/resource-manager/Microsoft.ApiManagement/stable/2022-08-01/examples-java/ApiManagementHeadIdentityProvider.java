import com.azure.resourcemanager.apimanagement.models.IdentityProviderType;

/** Samples for IdentityProvider GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementHeadIdentityProvider.json
     */
    /**
     * Sample code: ApiManagementHeadIdentityProvider.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadIdentityProvider(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .identityProviders()
            .getEntityTagWithResponse(
                "rg1", "apimService1", IdentityProviderType.AAD_B2C, com.azure.core.util.Context.NONE);
    }
}
