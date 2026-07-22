
/**
 * Samples for SharedGalleryInvites GallerySharingReject.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/sharedGalleryInviteExamples/SharedGalleryInvite_Reject.json
     */
    /**
     * Sample code: Reject a gallery shared to subscription.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void rejectAGallerySharedToSubscription(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getSharedGalleryInvites().gallerySharingReject("{location}",
            "480fd389-260b-41aa-bad9-0a83107c370c", "originalGalleryName", com.azure.core.util.Context.NONE);
    }
}
