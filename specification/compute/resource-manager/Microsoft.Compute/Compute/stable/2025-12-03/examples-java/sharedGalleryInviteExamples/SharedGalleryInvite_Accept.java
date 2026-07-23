
/**
 * Samples for SharedGalleryInvites GallerySharingAccept.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/sharedGalleryInviteExamples/SharedGalleryInvite_Accept.json
     */
    /**
     * Sample code: Accept a gallery shared to subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void acceptAGallerySharedToSubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSharedGalleryInvites().gallerySharingAccept("{location}",
            "480fd389-260b-41aa-bad9-0a83107c370c", "originalGalleryName", com.azure.core.util.Context.NONE);
    }
}
