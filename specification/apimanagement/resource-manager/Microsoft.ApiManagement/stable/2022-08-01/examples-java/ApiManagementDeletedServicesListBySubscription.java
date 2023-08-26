/** Samples for DeletedServices List. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementDeletedServicesListBySubscription.json
     */
    /**
     * Sample code: ApiManagementDeletedServicesListBySubscription.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeletedServicesListBySubscription(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.deletedServices().list(com.azure.core.util.Context.NONE);
    }
}
