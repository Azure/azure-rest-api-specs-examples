
/**
 * Samples for Triggers Get.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/imagebuilder/resource-manager/Microsoft.VirtualMachineImages/stable/2024-02-01/examples/GetTrigger.
     * json
     */
    /**
     * Sample code: Get a trigger resource.
     * 
     * @param manager Entry point to ImageBuilderManager.
     */
    public static void getATriggerResource(com.azure.resourcemanager.imagebuilder.ImageBuilderManager manager) {
        manager.triggers().getWithResponse("myResourceGroup", "myImageTemplate", "source",
            com.azure.core.util.Context.NONE);
    }
}
