import com.azure.resourcemanager.sphere.models.OSFeedType;
import com.azure.resourcemanager.sphere.models.UpdatePolicy;

/** Samples for DeviceGroups CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PutDeviceGroup.json
     */
    /**
     * Sample code: DeviceGroups_CreateOrUpdate.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void deviceGroupsCreateOrUpdate(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .deviceGroups()
            .define("MyDeviceGroup1")
            .withExistingProduct("MyResourceGroup1", "MyCatalog1", "MyProduct1")
            .withDescription("Description for MyDeviceGroup1")
            .withOsFeedType(OSFeedType.RETAIL)
            .withUpdatePolicy(UpdatePolicy.UPDATE_ALL)
            .create();
    }
}
