import com.azure.resourcemanager.azurestackhci.models.ExtendedLocation;
import com.azure.resourcemanager.azurestackhci.models.ExtendedLocationTypes;

/** Samples for VirtualHardDisksOperation CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azurestackhci/resource-manager/Microsoft.AzureStackHCI/preview/2023-09-01-preview/examples/PutVirtualHardDisk.json
     */
    /**
     * Sample code: PutVirtualHardDisk.
     *
     * @param manager Entry point to AzureStackHciManager.
     */
    public static void putVirtualHardDisk(com.azure.resourcemanager.azurestackhci.AzureStackHciManager manager) {
        manager
            .virtualHardDisksOperations()
            .define("test-vhd")
            .withRegion("West US2")
            .withExistingResourceGroup("test-rg")
            .withExtendedLocation(
                new ExtendedLocation()
                    .withName(
                        "/subscriptions/a95612cb-f1fa-4daa-a4fd-272844fa512c/resourceGroups/dogfoodarc/providers/Microsoft.ExtendedLocation/customLocations/dogfood-location")
                    .withType(ExtendedLocationTypes.CUSTOM_LOCATION))
            .withDiskSizeGB(32L)
            .create();
    }
}
