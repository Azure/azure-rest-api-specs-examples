import com.azure.resourcemanager.apimanagement.models.IdentityProviderType;

/** Samples for IdentityProvider Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeleteIdentityProvider.json
     */
    /**
     * Sample code: ApiManagementDeleteIdentityProvider.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteIdentityProvider(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .identityProviders()
            .deleteWithResponse("rg1", "apimService1", IdentityProviderType.AAD, "*", com.azure.core.util.Context.NONE);
    }
}
