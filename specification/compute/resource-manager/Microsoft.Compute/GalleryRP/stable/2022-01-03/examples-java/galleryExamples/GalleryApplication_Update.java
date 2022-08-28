import com.azure.core.util.Context;
import com.azure.resourcemanager.compute.models.GalleryApplicationUpdate;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/** Samples for GalleryApplications Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2022-01-03/examples/galleryExamples/GalleryApplication_Update.json
     */
    /**
     * Sample code: Update a simple gallery Application.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateASimpleGalleryApplication(com.azure.resourcemanager.AzureResourceManager azure) {
        azure
            .virtualMachines()
            .manager()
            .serviceClient()
            .getGalleryApplications()
            .update(
                "myResourceGroup",
                "myGalleryName",
                "myGalleryApplicationName",
                new GalleryApplicationUpdate()
                    .withDescription("This is the gallery application description.")
                    .withEula("This is the gallery application EULA.")
                    .withPrivacyStatementUri("myPrivacyStatementUri}")
                    .withReleaseNoteUri("myReleaseNoteUri")
                    .withSupportedOSType(OperatingSystemTypes.WINDOWS),
                Context.NONE);
    }
}
