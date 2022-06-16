import com.azure.core.util.Context;

/** Samples for VirtualMachineImageTemplates Cancel. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/CancelImageBuild.json
     */
    /**
     * Sample code: Cancel the image build based on the imageTemplate.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void cancelTheImageBuildBasedOnTheImageTemplate(
        com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().cancel("myResourceGroup", "myImageTemplate", Context.NONE);
    }
}
