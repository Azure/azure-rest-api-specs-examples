import com.azure.core.util.Context;

/** Samples for VirtualMachineImageTemplates GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2021-10-01/examples/GetImageTemplate.json
     */
    /**
     * Sample code: Retrieve an Image Template.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void retrieveAnImageTemplate(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager
            .virtualMachineImageTemplates()
            .getByResourceGroupWithResponse("myResourceGroup", "myImageTemplate", Context.NONE);
    }
}
