
/**
 * Samples for VirtualMachineImageTemplates ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/ListImageTemplatesByRg.json
     */
    /**
     * Sample code: List images by resource group.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void listImagesByResourceGroup(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
