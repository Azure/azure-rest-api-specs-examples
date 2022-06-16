import com.azure.core.util.Context;

/** Samples for Subscriptions ListLocations. */
public final class Main {
    /*
     * x-ms-original-file: specification/resources/resource-manager/Microsoft.Resources/stable/2021-01-01/examples/GetLocationsWithExtendedLocations.json
     */
    /**
     * Sample code: GetLocationsWithExtendedLocations.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getLocationsWithExtendedLocations(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .genericResources()
            .manager()
            .subscriptionClient()
            .getSubscriptions()
            .listLocations("291bba3f-e0a5-47bc-a099-3bdcb2a50a05", true, Context.NONE);
    }
}
