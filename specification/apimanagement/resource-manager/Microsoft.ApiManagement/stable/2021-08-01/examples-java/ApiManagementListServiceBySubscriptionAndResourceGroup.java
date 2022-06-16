import com.azure.core.util.Context;

/** Samples for ApiManagementService ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListServiceBySubscriptionAndResourceGroup.json
     */
    /**
     * Sample code: ApiManagementListServiceBySubscriptionAndResourceGroup.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListServiceBySubscriptionAndResourceGroup(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementServices().listByResourceGroup("rg1", Context.NONE);
    }
}
