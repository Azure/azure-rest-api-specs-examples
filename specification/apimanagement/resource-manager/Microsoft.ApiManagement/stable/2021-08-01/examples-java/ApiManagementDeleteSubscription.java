import com.azure.core.util.Context;

/** Samples for Subscription Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementDeleteSubscription.json
     */
    /**
     * Sample code: ApiManagementDeleteSubscription.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementDeleteSubscription(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.subscriptions().deleteWithResponse("rg1", "apimService1", "testsub", "*", Context.NONE);
    }
}
