import com.azure.core.util.Context;

/** Samples for AzureTrafficCollectorsBySubscription List. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkfunction/resource-manager/Microsoft.NetworkFunction/stable/2022-11-01/examples/AzureTrafficCollectorsBySubscriptionList.json
     */
    /**
     * Sample code: List of Traffic Collectors by Subscription.
     *
     * @param manager Entry point to AzureTrafficCollectorManager.
     */
    public static void listOfTrafficCollectorsBySubscription(
        com.azure.resourcemanager.networkfunction.AzureTrafficCollectorManager manager) {
        manager.azureTrafficCollectorsBySubscriptions().list(Context.NONE);
    }
}
