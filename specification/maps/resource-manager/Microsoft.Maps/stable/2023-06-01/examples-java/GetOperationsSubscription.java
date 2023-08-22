/** Samples for Maps List. */
public final class Main {
    /*
     * x-ms-original-file: specification/maps/resource-manager/Microsoft.Maps/stable/2023-06-01/examples/GetOperationsSubscription.json
     */
    /**
     * Sample code: Get Operations by Subscription.
     *
     * @param manager Entry point to AzureMapsManager.
     */
    public static void getOperationsBySubscription(com.azure.resourcemanager.maps.AzureMapsManager manager) {
        manager.maps().list(com.azure.core.util.Context.NONE);
    }
}
