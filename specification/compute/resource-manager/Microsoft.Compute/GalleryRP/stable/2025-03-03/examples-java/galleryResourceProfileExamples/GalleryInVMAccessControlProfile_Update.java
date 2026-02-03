
import com.azure.resourcemanager.compute.models.EndpointTypes;
import com.azure.resourcemanager.compute.models.GalleryInVMAccessControlProfileProperties;
import com.azure.resourcemanager.compute.models.GalleryInVMAccessControlProfileUpdate;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/**
 * Samples for GalleryInVMAccessControlProfiles Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2025-03-03/examples/
     * galleryResourceProfileExamples/GalleryInVMAccessControlProfile_Update.json
     */
    /**
     * Sample code: Update a gallery inVMAccessControlProfile.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateAGalleryInVMAccessControlProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryInVMAccessControlProfiles().update(
            "myResourceGroup", "myGalleryName", "myInVMAccessControlProfileName",
            new GalleryInVMAccessControlProfileUpdate().withProperties(new GalleryInVMAccessControlProfileProperties()
                .withOsType(OperatingSystemTypes.LINUX).withApplicableHostEndpoint(EndpointTypes.WIRE_SERVER)),
            com.azure.core.util.Context.NONE);
    }
}
