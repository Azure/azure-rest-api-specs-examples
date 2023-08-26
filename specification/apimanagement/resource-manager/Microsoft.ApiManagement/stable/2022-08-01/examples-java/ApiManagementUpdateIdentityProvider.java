import com.azure.resourcemanager.apimanagement.models.IdentityProviderContract;
import com.azure.resourcemanager.apimanagement.models.IdentityProviderType;

/** Samples for IdentityProvider Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateIdentityProvider.json
     */
    /**
     * Sample code: ApiManagementUpdateIdentityProvider.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateIdentityProvider(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        IdentityProviderContract resource =
            manager
                .identityProviders()
                .getWithResponse("rg1", "apimService1", IdentityProviderType.FACEBOOK, com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withClientId("updatedfacebookid")
            .withClientSecret("updatedfacebooksecret")
            .withIfMatch("*")
            .apply();
    }
}
