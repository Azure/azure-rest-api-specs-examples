import com.azure.resourcemanager.sphere.models.CapabilityType;
import com.azure.resourcemanager.sphere.models.GenerateCapabilityImageRequest;
import java.util.Arrays;

/** Samples for Devices GenerateCapabilityImage. */
public final class Main {
    /*
     * x-ms-original-file: specification/sphere/resource-manager/Microsoft.AzureSphere/preview/2022-09-01-preview/examples/PostGenerateDeviceCapabilityImage.json
     */
    /**
     * Sample code: Devices_GenerateCapabilityImage.
     *
     * @param manager Entry point to AzureSphereManager.
     */
    public static void devicesGenerateCapabilityImage(com.azure.resourcemanager.sphere.AzureSphereManager manager) {
        manager
            .devices()
            .generateCapabilityImage(
                "MyResourceGroup1",
                "MyCatalog1",
                "MyProduct1",
                "myDeviceGroup1",
                "00000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000",
                new GenerateCapabilityImageRequest()
                    .withCapabilities(Arrays.asList(CapabilityType.APPLICATION_DEVELOPMENT)),
                com.azure.core.util.Context.NONE);
    }
}
