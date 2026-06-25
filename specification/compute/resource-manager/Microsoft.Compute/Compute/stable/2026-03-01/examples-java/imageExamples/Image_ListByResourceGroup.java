
/**
 * Samples for Images ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-01/imageExamples/Image_ListByResourceGroup.json
     */
    /**
     * Sample code: List all virtual machine images in a resource group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listAllVirtualMachineImagesInAResourceGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getImages().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
