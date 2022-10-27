import com.azure.core.util.Context;

/** Samples for Profiles List. */
public final class Main {
    /*
     * x-ms-original-file: specification/trafficmanager/resource-manager/Microsoft.Network/stable/2018-04-01/examples/Profile-GET-BySubscription.json
     */
    /**
     * Sample code: ListBySubscription.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listBySubscription(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.trafficManagerProfiles().manager().serviceClient().getProfiles().list(Context.NONE);
    }
}
