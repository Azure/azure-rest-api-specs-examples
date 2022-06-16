import com.azure.core.util.Context;

/** Samples for VirtualMachineImageTemplates Run. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/RunImageTemplate.json
     */
    /**
     * Sample code: Create image(s) from existing imageTemplate.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void createImageSFromExistingImageTemplate(
        com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().run("myResourceGroup", "myImageTemplate", Context.NONE);
    }
}
