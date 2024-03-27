
import com.azure.resourcemanager.sphere.models.ClaimDevicesRequest;
import java.util.Arrays;

/**
 * Samples for DeviceGroups ClaimDevices.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PostClaimDevices.json
     */
    /**
     * Sample code: DeviceGroups_ClaimDevices.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void deviceGroupsClaimDevices(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager.deviceGroups().claimDevices("MyResourceGroup1", "MyCatalog1", "MyProduct1", "MyDeviceGroup1",
            new ClaimDevicesRequest().withDeviceIdentifiers(Arrays.asList(
                "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")),
            com.azure.core.util.Context.NONE);
    }
}
