import com.azure.core.management.SubResource;
import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.fluent.models.ImageInner;
import com.azure.resourcemanager.compute.models.DiskEncryptionSetParameters;
import com.azure.resourcemanager.compute.models.ImageOSDisk;
import com.azure.resourcemanager.compute.models.ImageStorageProfile;
import com.azure.resourcemanager.compute.models.OperatingSystemStateTypes;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/** Samples for Images CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/stable/2021-07-01/examples/compute/CreateAnImageFromAManagedDiskWithDiskEncryptionSet.json
     */
    /**
     * Sample code: Create a virtual machine image from a managed disk with DiskEncryptionSet resource.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAVirtualMachineImageFromAManagedDiskWithDiskEncryptionSetResource(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getImages()
            .createOrUpdate(
                "myResourceGroup",
                "myImage",
                new ImageInner()
                    .withLocation("West US")
                    .withStorageProfile(
                        new ImageStorageProfile()
                            .withOsDisk(
                                new ImageOSDisk()
                                    .withManagedDisk(
                                        new SubResource()
                                            .withId(
                                                "subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/disks/myManagedDisk"))
                                    .withDiskEncryptionSet(
                                        new DiskEncryptionSetParameters()
                                            .withId(
                                                "/subscriptions/{subscription-id}/resourceGroups/myResourceGroup/providers/Microsoft.Compute/diskEncryptionSets/{existing-diskEncryptionSet-name}"))
                                    .withOsType(OperatingSystemTypes.LINUX)
                                    .withOsState(OperatingSystemStateTypes.GENERALIZED))),
                Context.NONE);
    }
}
