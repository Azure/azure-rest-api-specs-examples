import com.azure.core.util.Context;

/** Samples for DelegationSettings GetEntityTag. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementHeadDelegationSettings.json
     */
    /**
     * Sample code: ApiManagementHeadDelegationSettings.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementHeadDelegationSettings(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.delegationSettings().getEntityTagWithResponse("rg1", "apimService1", Context.NONE);
    }
}
