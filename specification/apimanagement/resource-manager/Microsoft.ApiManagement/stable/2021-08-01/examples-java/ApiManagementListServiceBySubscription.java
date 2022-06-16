import com.azure.core.util.Context;

/** Samples for ApiManagementService List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementListServiceBySubscription.json
     */
    /**
     * Sample code: ApiManagementListServiceBySubscription.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementListServiceBySubscription(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.apiManagementServices().list(Context.NONE);
    }
}
