
/**
 * Samples for VirtualMachineImageTemplates Run.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2024-02-01/examples/
     * RunImageTemplate.json
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
