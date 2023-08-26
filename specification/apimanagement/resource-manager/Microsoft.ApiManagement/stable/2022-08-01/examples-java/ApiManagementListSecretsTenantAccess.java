import com.azure.resourcemanager.apimanagement.models.AccessIdName;

/** Samples for TenantAccess ListSecrets. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementListSecretsTenantAccess.json
     */
    /**
     * Sample code: ApiManagementListSecretsTenantAccess.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListSecretsTenantAccess(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .tenantAccess()
            .listSecretsWithResponse("rg1", "apimService1", AccessIdName.ACCESS, com.azure.core.util.Context.NONE);
    }
}
