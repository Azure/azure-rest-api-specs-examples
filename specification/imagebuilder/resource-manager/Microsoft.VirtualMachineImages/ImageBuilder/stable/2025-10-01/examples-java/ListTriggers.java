
/**
 * Samples for Triggers ListByImageTemplate.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/ListTriggers.json
     */
    /**
     * Sample code: List triggers by image template.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void listTriggersByImageTemplate(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.triggers().listByImageTemplate("myResourceGroup", "myImageTemplate", com.azure.core.util.Context.NONE);
    }
}
