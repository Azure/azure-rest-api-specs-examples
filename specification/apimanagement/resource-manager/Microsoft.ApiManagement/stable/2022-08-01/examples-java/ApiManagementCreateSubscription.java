import com.azure.resourcemanager.apimanagement.models.SubscriptionCreateParameters;

/** Samples for Subscription CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementCreateSubscription.json
     */
    /**
     * Sample code: ApiManagementCreateSubscription.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementCreateSubscription(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .subscriptions()
            .createOrUpdateWithResponse(
                "rg1",
                "apimService1",
                "testsub",
                new SubscriptionCreateParameters()
                    .withOwnerId(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/users/57127d485157a511ace86ae7")
                    .withScope(
                        "/subscriptions/subid/resourceGroups/rg1/providers/Microsoft.ApiManagement/service/apimService1/products/5600b59475ff190048060002")
                    .withDisplayName("testsub"),
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
