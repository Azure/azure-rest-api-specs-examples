
/**
 * Samples for VirtualMachineImageTemplates GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2023-07-01/examples/
     * GetImageTemplate.json
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
