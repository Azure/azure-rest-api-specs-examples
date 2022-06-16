import com.azure.core.util.Context;

/** Samples for IdentityProvider ListByService. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListIdentityProviders.json
     */
    /**
     * Sample code: ApiManagementListIdentityProviders.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListIdentityProviders(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.identityProviders().listByService("rg1", "apimService1", Context.NONE);
    }
}
