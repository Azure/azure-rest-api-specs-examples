import com.azure.core.util.Context;

/** Samples for DelegationSettings ListSecrets. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListSecretsPortalSettingsValidationKey.json
     */
    /**
     * Sample code: ApiManagementListSecretsPortalSettings.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListSecretsPortalSettings(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.delegationSettings().listSecretsWithResponse("rg1", "apimService1", Context.NONE);
    }
}
