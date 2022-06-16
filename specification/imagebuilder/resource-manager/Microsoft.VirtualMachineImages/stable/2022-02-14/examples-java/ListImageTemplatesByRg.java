import com.azure.core.util.Context;

/** Samples for VirtualMachineImageTemplates ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/ListImageTemplatesByRg.json
     */
    /**
     * Sample code: List images by resource group.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void listImagesByResourceGroup(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().listByResourceGroup("myResourceGroup", Context.NONE);
    }
}
