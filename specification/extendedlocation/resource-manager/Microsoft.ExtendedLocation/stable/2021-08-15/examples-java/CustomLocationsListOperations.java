import com.azure.core.util.Context;

/** Samples for CustomLocations ListOperations. */
public final class Main {
    /*
     * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsListOperations.json
     */
    /**
     * Sample code: List Custom Locations operations.
     *
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void listCustomLocationsOperations(
        com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager.customLocations().listOperations(Context.NONE);
    }
}
