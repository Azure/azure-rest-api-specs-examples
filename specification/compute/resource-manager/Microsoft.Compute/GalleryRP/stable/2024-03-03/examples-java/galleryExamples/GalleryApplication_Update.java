
import com.azure.resourcemanager.compute.models.GalleryApplicationCustomAction;
import com.azure.resourcemanager.compute.models.GalleryApplicationCustomActionParameter;
import com.azure.resourcemanager.compute.models.GalleryApplicationCustomActionParameterType;
import com.azure.resourcemanager.compute.models.GalleryApplicationUpdate;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;
import java.util.Arrays;

/**
 * Samples for GalleryApplications Update.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/compute/resource-manager/Microsoft.Compute/GalleryRP/stable/2024-03-03/examples/galleryExamples/
     * GalleryApplication_Update.json
     */
    /**
     * Sample code: Update a simple gallery Application.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void updateASimpleGalleryApplication(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.virtualMachines().manager().serviceClient().getGalleryApplications()
            .update("myResourceGroup", "myGalleryName", "myGalleryApplicationName",
                new GalleryApplicationUpdate().withDescription("This is the gallery application description.")
                    .withEula("This is the gallery application EULA.").withPrivacyStatementUri("myPrivacyStatementUri}")
                    .withReleaseNoteUri("myReleaseNoteUri").withSupportedOSType(OperatingSystemTypes.WINDOWS)
                    .withCustomActions(Arrays.asList(new GalleryApplicationCustomAction().withName("myCustomAction")
                        .withScript("myCustomActionScript").withDescription("This is the custom action description.")
                        .withParameters(Arrays
                            .asList(new GalleryApplicationCustomActionParameter().withName("myCustomActionParameter")
                                .withRequired(false).withType(GalleryApplicationCustomActionParameterType.STRING)
                                .withDefaultValue("default value of parameter.")
                                .withDescription("This is the description of the parameter"))))),
                com.azure.core.util.Context.NONE);
    }
}
