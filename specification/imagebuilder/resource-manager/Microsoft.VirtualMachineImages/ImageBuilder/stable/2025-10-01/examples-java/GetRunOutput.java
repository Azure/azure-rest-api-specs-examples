
/**
 * Samples for VirtualMachineImageTemplates GetRunOutput.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-10-01/GetRunOutput.json
     */
    /**
     * Sample code: Retrieve single runOutput.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void retrieveSingleRunOutput(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().getRunOutputWithResponse("myResourceGroup", "myImageTemplate",
            "myManagedImageOutput", com.azure.core.util.Context.NONE);
    }
}
