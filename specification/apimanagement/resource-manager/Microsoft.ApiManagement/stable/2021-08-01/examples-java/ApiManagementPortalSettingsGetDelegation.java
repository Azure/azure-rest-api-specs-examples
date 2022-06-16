import com.azure.core.util.Context;

/** Samples for DelegationSettings Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementPortalSettingsGetDelegation.json
     */
    /**
     * Sample code: ApiManagementPortalSettingsGetDelegation.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementPortalSettingsGetDelegation(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.delegationSettings().getWithResponse("rg1", "apimService1", Context.NONE);
    }
}
