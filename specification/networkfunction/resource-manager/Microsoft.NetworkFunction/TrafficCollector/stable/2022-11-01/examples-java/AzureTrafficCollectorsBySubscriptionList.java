
/**
 * Samples for AzureTrafficCollectorsBySubscription List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2022-11-01/AzureTrafficCollectorsBySubscriptionList.json
     */
    /**
     * Sample code: List of Traffic Collectors by Subscription.
     * 
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void listOfTrafficCollectorsBySubscription(
        com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.azureTrafficCollectorsBySubscriptions().list(com.azure.core.util.Context.NONE);
    }
}
