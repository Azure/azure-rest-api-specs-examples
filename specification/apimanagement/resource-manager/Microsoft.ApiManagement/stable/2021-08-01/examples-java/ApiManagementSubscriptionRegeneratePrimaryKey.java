import com.azure.core.util.Context;

/** Samples for Subscription RegeneratePrimaryKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementSubscriptionRegeneratePrimaryKey.json
     */
    /**
     * Sample code: ApiManagementSubscriptionRegeneratePrimaryKey.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementSubscriptionRegeneratePrimaryKey(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.subscriptions().regeneratePrimaryKeyWithResponse("rg1", "apimService1", "testsub", Context.NONE);
    }
}
