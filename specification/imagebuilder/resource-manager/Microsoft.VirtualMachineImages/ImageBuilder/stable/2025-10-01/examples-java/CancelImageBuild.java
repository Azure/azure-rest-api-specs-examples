
/**
 * Samples for VirtualMachineImageTemplates Cancel.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/CancelImageBuild.json
     */
    /**
     * Sample code: Cancel the image build based on the imageTemplate.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void
        cancelTheImageBuildBasedOnTheImageTemplate(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().cancel("myResourceGroup", "myImageTemplate",
            com.azure.core.util.Context.NONE);
    }
}
