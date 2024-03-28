
import com.azure.resourcemanager.sphere.models.Device;

/**
 * Samples for Devices Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sphere/resource-manager/Microsoft.AzureSphere/stable/2024-04-01/examples/PatchDevice.json
     */
    /**
     * Sample code: Devices_Update.
     * 
     * @param manager Entry point to AzureSphereManager.
     */
    public static void devicesUpdate(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        Device resource = manager.devices().getWithResponse("MyResourceGroup1", "MyCatalog1", "MyProduct1",
            "MyDeviceGroup1",
            "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
            com.azure.core.util.Context.NONE).getValue();
        resource.update().apply();
    }
}
