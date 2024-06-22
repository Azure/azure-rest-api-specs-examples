
/**
 * Samples for Triggers Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2024-02-01/examples/
     * DeleteTrigger.json
     */
    /**
     * Sample code: Delete a trigger resource.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void deleteATriggerResource(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.triggers().delete("myResourceGroup", "myImageTemplate", "trigger1", com.azure.core.util.Context.NONE);
    }
}
