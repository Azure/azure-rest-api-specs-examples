import com.azure.core.util.Context;

/** Samples for CustomLocations ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/extendedlocation/resource-manager/Microsoft.ExtendedLocation/stable/2021-08-15/examples/CustomLocationsListByResourceGroup.json
     */
    /**
     * Sample code: List Custom Locations by resource group.
     *
     * @param manager Entry point to CustomLocationsManager.
     */
    public static void listCustomLocationsByResourceGroup(
        com.azure.resourcemanager.extendedlocation.CustomLocationsManager manager) {
        manager.customLocations().listByResourceGroup("testresourcegroup", Context.NONE);
    }
}
