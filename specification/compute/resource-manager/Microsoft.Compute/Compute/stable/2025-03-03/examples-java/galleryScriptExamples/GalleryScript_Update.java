
import com.azure.resourcemanager.compute.models.GalleryScriptUpdate;
import com.azure.resourcemanager.compute.models.OperatingSystemTypes;

/**
 * Samples for GalleryScripts Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-03/galleryScriptExamples/GalleryScript_Update.json
     */
    /**
     * Sample code: Update a simple gallery Script.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void updateASimpleGalleryScript(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleryScripts().update("myResourceGroup", "myGalleryName", "myGalleryScriptName",
            new GalleryScriptUpdate().withDescription("This is the gallery script description.")
                .withEula("This is the gallery script EULA.").withPrivacyStatementUri("{myPrivacyStatementUri}")
                .withReleaseNoteUri("{myReleaseNoteUri}").withSupportedOSType(OperatingSystemTypes.WINDOWS),
            com.azure.core.util.Context.NONE);
    }
}
