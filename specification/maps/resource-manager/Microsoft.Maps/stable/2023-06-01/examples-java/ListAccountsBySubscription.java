/** Samples for Accounts List. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/ListAccountsBySubscription.json
     */
    /**
     * Sample code: List Accounts By Subscription.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void listAccountsBySubscription(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager.accounts().list(com.azure.core.util.Context.NONE);
    }
}
