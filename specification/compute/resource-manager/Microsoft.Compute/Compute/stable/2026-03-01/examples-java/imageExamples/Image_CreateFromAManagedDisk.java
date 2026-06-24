
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.compute.fluent.models.ImageInner;
import com.azure.resourcemanager.compute.models.ImageOSDisk;
import com.azure.resourcemanager.compute.models.ImageStorageProfile;
import com.azure.resourcemanager.compute.models.OperatingSystemStateTypes;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/**
 * Samples for Images CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/imageExamples/Image_CreateFromAManagedDisk.json
     */
    /**
     * Sample code: Create a virtual machine image from a managed disk.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        createAVirtualMachineImageFromAManagedDisk(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getImages().createOrUpdate("myResourceGroup", "myImage",
            new ImageInner().withLocation("West US").withStorageProfile(new ImageStorageProfile()
                .withOsDisk(new ImageOSDisk().withManagedDisk(new SubResource().withId(
                    "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"))
                    .withOsType(OperatingSystemTypes.LINUX).withOsState(OperatingSystemStateTypes.GENERALIZED))
                .withZoneResilient(true)),
            com.azure.core.util.Context.NONE);
    }
}
