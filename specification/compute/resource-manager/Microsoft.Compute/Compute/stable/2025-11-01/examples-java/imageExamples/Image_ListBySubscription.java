
/**
 * Samples for Images List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/imageExamples/Image_ListBySubscription.json
     */
    /**
     * Sample code: List all virtual machine images in a subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void
        listAllVirtualMachineImagesInASubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getImages().list(com.azure.core.util.Context.NONE);
    }
}
