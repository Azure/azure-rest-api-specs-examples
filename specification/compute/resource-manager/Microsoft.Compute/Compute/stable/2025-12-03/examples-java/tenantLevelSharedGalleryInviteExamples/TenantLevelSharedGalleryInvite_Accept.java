
/**
 * Samples for TenantLevelSharedGalleryInvites TenantLevelGallerySharingAccept.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-12-03/tenantLevelSharedGalleryInviteExamples/TenantLevelSharedGalleryInvite_Accept.json
     */
    /**
     * Sample code: Accept a gallery shared to tenant.
     * 
     * @param manager Entry point to ComputeManager.
     */
    public static void acceptAGallerySharedToTenant(com.azure.resourcemanager.compute.ComputeManager manager) {
        manager.serviceClient().getTenantLevelSharedGalleryInvites().tenantLevelGallerySharingAccept("{location}",
            "480fd389-260b-41aa-bad9-0a83107c370c", "originalGalleryName", com.azure.core.util.Context.NONE);
    }
}
