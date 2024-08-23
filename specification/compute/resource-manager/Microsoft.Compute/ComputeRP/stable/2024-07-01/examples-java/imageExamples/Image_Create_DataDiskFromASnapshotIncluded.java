
import com.azure.core.management.SubResource;
import com.azure.resourcemanager.compute.fluent.models.ImageInner;
import com.azure.resourcemanager.compute.models.ImageDataDisk;
import com.azure.resourcemanager.compute.models.ImageOSDisk;
import com.azure.resourcemanager.compute.models.ImageStorageProfile;
import com.azure.resourcemanager.compute.models.OperatingSystemStateTypes;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;
import java.util.Arrays;

/**
 * Samples for Images CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/ComputeRP/stable/2024-07-01/examples/imageExamples/
     * Image_Create_DataDiskFromASnapshotIncluded.json
     */
    /**
     * Sample code: Create a virtual machine image that includes a data disk from a snapshot.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAVirtualMachineImageThatIncludesADataDiskFromASnapshot(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getImages().createOrUpdate("myResourceGroup", "myImage",
            new ImageInner().withLocation("West US").withStorageProfile(new ImageStorageProfile()
                .withOsDisk(new ImageOSDisk().withSnapshot(new SubResource().withId(
                    "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot"))
                    .withOsType(OperatingSystemTypes.LINUX).withOsState(OperatingSystemStateTypes.GENERALIZED))
                .withDataDisks(Arrays.asList(new ImageDataDisk().withSnapshot(new SubResource().withId(
                    "subscriptions/{subscriptionId}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/snapshots/mySnapshot2"))
                    .withLun(1)))
                .withZoneResilient(true)),
            com.azure.core.util.Context.NONE);
    }
}
