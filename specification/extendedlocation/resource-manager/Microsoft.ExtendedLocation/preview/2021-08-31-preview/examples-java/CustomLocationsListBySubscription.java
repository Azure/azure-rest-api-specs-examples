import com.azure.core.util.Context;

/** Samples for CustomLocations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/CustomLocationsListBySubscription.json
     */
    /**
     * Sample code: List Custom Locations by subscription.
     *
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void listCustomLocationsBySubscription(
        com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager.customLocations().list(Context.NONE);
    }
}
