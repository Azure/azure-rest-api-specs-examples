import com.azure.core.util.Context;

/** Samples for CustomLocations ListEnabledResourceTypes. */
public final class Main {
    /*
     * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/preview/2021-08-31-preview/examples/CustomLocationsListEnabledResourceTypes.json
     */
    /**
     * Sample code: Get Custom Location.
     *
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void getCustomLocation(com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager.customLocations().listEnabledResourceTypes("testresourcegroup", "customLocation01", Context.NONE);
    }
}
