
/**
 * Samples for VirtualMachineImageTemplates ListRunOutputs.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2024-02-01/examples/
     * ListRunOutputs.json
     */
    /**
     * Sample code: Retrieve a list of all outputs created by the last run of an Image Template.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void retrieveAListOfAllOutputsCreatedByTheLastRunOfAnImageTemplate(
        com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.virtualMachineImageTemplates().listRunOutputs("myResourceGroup", "myImageTemplate",
            com.azure.core.util.Context.NONE);
    }
}
