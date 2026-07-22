
/**
 * Samples for TenantLevelSharedGalleryInvites TenantLevelGallerySharingReject.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/tenantLevelSharedGalleryInviteExamples/TenantLevelSharedGalleryInvite_Reject.json
     */
    /**
     * Sample code: Reject a gallery shared to tenant.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void rejectAGallerySharedToTenant(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getTenantLevelSharedGalleryInvites().tenantLevelGallerySharingReject("{location}",
            "480fd389-260b-41aa-bad9-0a83107c370c", "originalGalleryName", com.azure.core.util.Context.NONE);
    }
}
