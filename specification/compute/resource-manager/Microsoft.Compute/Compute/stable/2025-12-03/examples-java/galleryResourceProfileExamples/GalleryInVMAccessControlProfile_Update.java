
import com.azure.resourcemanager.compute.models.EndpointTypes;
import com.azure.resourcemanager.compute.models.GalleryInVMAccessControlProfileProperties;
import com.azure.resourcemanager.compute.models.GalleryInVMAccessControlProfileUpdate;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/**
 * Samples for GalleryInVMAccessControlProfiles Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryResourceProfileExamples/GalleryInVMAccessControlProfile_Update.json
     */
    /**
     * Sample code: Update a gallery inVMAccessControlProfile.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        updateAGalleryInVMAccessControlProfile(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryInVMAccessControlProfiles().update("myResourceGroup", "myGalleryName",
            "myInVMAccessControlProfileName",
            new GalleryInVMAccessControlProfileUpdate().withProperties(new GalleryInVMAccessControlProfileProperties()
                .withOsType(OperatingSystemTypes.LINUX).withApplicableHostEndpoint(EndpointTypes.WIRE_SERVER)),
            com.azure.core.util.Context.NONE);
    }
}
