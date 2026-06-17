
/**
 * Samples for VirtualMachineImageTemplates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/DeleteImageTemplate.json
     */
    /**
     * Sample code: Delete an Image Template.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void deleteAnImageTemplate(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().delete("myResourceGroup", "myImageTemplate",
            com.azure.core.util.Context.NONE);
    }
}
