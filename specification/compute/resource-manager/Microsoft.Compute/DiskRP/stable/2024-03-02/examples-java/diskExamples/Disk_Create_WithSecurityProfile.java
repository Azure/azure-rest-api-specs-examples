
import com.azure.resourcemanager.compute.fluent.models.DiskInner;
import com.azure.resourcemanager.compute.models.CreationData;
import com.azure.resourcemanager.compute.models.DiskCreateOption;
import com.azure.resourcemanager.compute.models.DiskSecurityProfile;
import com.azure.resourcemanager.compute.models.DiskSecurityTypes;
import com.azure.resourcemanager.compute.models.ImageDiskReference;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/**
 * Samples for Disks CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/DiskRP/stable/2024-03-02/examples/diskExamples/
     * Disk_Create_WithSecurityProfile.json
     */
    /**
     * Sample code: Create a managed disk with security profile.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void createAManagedDiskWithSecurityProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getDisks().createOrUpdate("myResourceGroup", "myDisk",
            new DiskInner().withLocation("North Central US").withOsType(OperatingSystemTypes.WINDOWS)
                .withCreationData(new CreationData().withCreateOption(DiskCreateOption.FROM_IMAGE)
                    .withImageReference(new ImageDiskReference().withId(
                        "/Subscriptions/{subscriptionId}/Providers/Microsoft.Compute/Locations/uswest/Publishers/Microsoft/ArtifactTypes/VMImage/Offers/{offer}")))
                .withSecurityProfile(new DiskSecurityProfile().withSecurityType(DiskSecurityTypes.TRUSTED_LAUNCH)),
            com.azure.core.util.Context.NONE);
    }
}
