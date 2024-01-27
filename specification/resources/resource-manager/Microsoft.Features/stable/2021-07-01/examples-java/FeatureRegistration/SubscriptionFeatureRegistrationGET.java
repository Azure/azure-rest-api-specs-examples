
/** Samples for SubscriptionFeatureRegistrations Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/resources/resource-manager/Microsoft.Features/stable/2021-07-01/examples/FeatureRegistration/
     * SubscriptionFeatureRegistrationGET.json
     */
    /**
     * Sample code: Gets a feature registration.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getsAFeatureRegistration(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.genericResources().manager().featureClient().getSubscriptionFeatureRegistrations().getWithResponse(
            "subscriptionFeatureRegistrationGroupTestRG", "testFeature", com.azure.core.util.Context.NONE);
    }
}
