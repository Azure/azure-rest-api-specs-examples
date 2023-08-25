import com.azure.resourcemanager.apimanagement.models.SubscriptionUpdateParameters;

/** Samples for Subscription Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateSubscription.json
     */
    /**
     * Sample code: ApiManagementUpdateSubscription.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateSubscription(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .subscriptions()
            .updateWithResponse(
                "rg1",
                "apimService1",
                "testsub",
                "*",
                new SubscriptionUpdateParameters().withDisplayName("testsub"),
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
