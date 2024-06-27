
/**
 * Samples for SubscriptionFeatureRegistrations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/FeatureRegistration/
     * SubscriptionFeatureRegistrationDELETE.json
     */
    /**
     * Sample code: Deletes a feature registration.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void deletesAFeatureRegistration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().featureClient().getSubscriptionFeatureRegistrations().deleteWithResponse(
            "subscriptionFeatureRegistrationGroupTestRG", "testFeature", com.azure.core.util.Context.NONE);
    }
}
