
/**
 * Samples for VirtualMachineImageTemplates GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/GetImageTemplate.json
     */
    /**
     * Sample code: Retrieve an Image Template.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void retrieveAnImageTemplate(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().getByResourceGroupWithResponse("myResourceGroup", "myImageTemplate",
            com.azure.core.util.Context.NONE);
    }
}
