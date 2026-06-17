
import com.azure.resourcemanager.imagebuilder.models.SourceImageTriggerProperties;

/**
 * Samples for Triggers CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/CreateSourceImageTrigger.json
     */
    /**
     * Sample code: Create or update a source image type trigger.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void
        createOrUpdateASourceImageTypeTrigger(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.triggers().define("source").withExistingImageTemplate("myResourceGroup", "myImageTemplate")
            .withProperties(new SourceImageTriggerProperties()).create();
    }
}
