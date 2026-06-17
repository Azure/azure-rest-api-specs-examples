
/**
 * Samples for VirtualMachineImageTemplates Run.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/RunImageTemplate.json
     */
    /**
     * Sample code: Create image(s) from existing imageTemplate.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void
        createImageSFromExistingImageTemplate(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().run("myResourceGroup", "myImageTemplate",
            com.azure.core.util.Context.NONE);
    }
}
