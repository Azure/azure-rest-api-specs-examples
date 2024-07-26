
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.ImageDiskReference;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/
     * Disk_Create_FromAnAzureComputeGalleryCommunityImage.json
     */
    /**
     * Sample code: Create a managed disk from an Azure Compute Gallery community image.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAManagedDiskFromAnAzureComputeGalleryCommunityImage(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk",
            new DiskInner().withLocation("West US").withOsType(OperatingSystemTypes.WINDOWS)
                .withCreationData(new CreationData().withCreateOption(DiskCreateOption.FROM_IMAGE)
                    .withGalleryImageReference(new ImageDiskReference().withCommunityGalleryImageId(
                        "/CommunityGalleries/{communityGalleryPublicGalleryName}/Images/{imageName}/Versions/1.0.0"))),
            com.azure.core.util.Context.NONE);
    }
}
