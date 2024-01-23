
/**
 * Samples for VirtualMachineImageTemplates Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2023-07-01/examples/
     * DeleteImageTemplate.json
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
