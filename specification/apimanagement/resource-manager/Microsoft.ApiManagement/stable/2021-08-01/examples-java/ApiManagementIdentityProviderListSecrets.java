import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.IdentityProviderType;

/** Samples for IdentityProvider ListSecrets. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementIdentityProviderListSecrets.json
     */
    /**
     * Sample code: ApiManagementIdentityProviderListSecrets.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementIdentityProviderListSecrets(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .identityProviders()
            .listSecretsWithResponse("rg1", "apimService1", IdentityProviderType.AAD_B2C, Context.NONE);
    }
}
