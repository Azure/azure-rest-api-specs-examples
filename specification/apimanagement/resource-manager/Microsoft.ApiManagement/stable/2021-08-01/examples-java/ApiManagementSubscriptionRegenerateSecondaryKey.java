import com.azure.core.util.Context;

/** Samples for Subscription RegenerateSecondaryKey. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementSubscriptionRegenerateSecondaryKey.json
     */
    /**
     * Sample code: ApiManagementSubscriptionRegenerateSecondaryKey.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementSubscriptionRegenerateSecondaryKey(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager.subscriptions().regenerateSecondaryKeyWithResponse("rg1", "apimService1", "testsub", Context.NONE);
    }
}
