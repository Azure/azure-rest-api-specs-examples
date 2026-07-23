
/**
 * Samples for Galleries ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/galleryExamples/Gallery_ListByResourceGroup.json
     */
    /**
     * Sample code: List galleries in a resource group.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void listGalleriesInAResourceGroup(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getGalleries().listByResourceGroup("myResourceGroup", com.azure.core.util.Context.NONE);
    }
}
