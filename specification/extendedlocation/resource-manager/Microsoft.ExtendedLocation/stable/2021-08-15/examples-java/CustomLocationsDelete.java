import com.azure.core.util.Context;

/** Samples for CustomLocations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsDelete.json
     */
    /**
     * Sample code: Delete Custom Location.
     *
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void deleteCustomLocation(com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager.customLocations().delete("testresourcegroup", "customLocation01", Context.NONE);
    }
}
