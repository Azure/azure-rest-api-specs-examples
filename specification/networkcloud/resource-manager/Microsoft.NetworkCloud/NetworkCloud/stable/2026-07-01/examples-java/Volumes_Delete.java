
/**
 * Samples for Volumes Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-07-01/Volumes_Delete.json
     */
    /**
     * Sample code: Delete volume.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteVolume(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.volumes().delete("resourceGroupName", "volumeName", null, null, com.azure.core.util.Context.NONE);
    }
}
