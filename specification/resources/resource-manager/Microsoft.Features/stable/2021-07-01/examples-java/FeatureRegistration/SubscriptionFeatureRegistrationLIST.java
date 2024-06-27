
/**
 * Samples for SubscriptionFeatureRegistrations ListBySubscription.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/FeatureRegistration/
     * SubscriptionFeatureRegistrationLIST.json
     */
    /**
     * Sample code: Gets a list of feature registrations.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAListOfFeatureRegistrations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().featureClient().getSubscriptionFeatureRegistrations()
            .listBySubscription("subscriptionFeatureRegistrationGroupTestRG", com.azure.core.util.Context.NONE);
    }
}
