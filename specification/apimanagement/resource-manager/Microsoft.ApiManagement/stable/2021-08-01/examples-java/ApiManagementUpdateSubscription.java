import com.azure.core.util.Context;
import com.azure.resourcemanager.apimanagement.models.SubscriptionUpdateParameters;

/** Samples for Subscription Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2021-08-01/examples/ApiManagementUpdateSubscription.json
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
                Context.NONE);
    }
}
