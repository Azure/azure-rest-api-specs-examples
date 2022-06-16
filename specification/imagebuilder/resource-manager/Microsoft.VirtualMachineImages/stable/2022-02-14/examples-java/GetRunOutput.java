import com.azure.core.util.Context;

/** Samples for VirtualMachineImageTemplates GetRunOutput. */
public final class Main {
    /*
     * x-ms-original-file: specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2022-02-14/examples/GetRunOutput.json
     */
    /**
     * Sample code: Retrieve single runOutput.
     *
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void retrieveSingleRunOutput(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager
            .virtualMachineImageTemplates()
            .getRunOutputWithResponse("myResourceGroup", "myImageTemplate", "myManagedImageOutput", Context.NONE);
    }
}
