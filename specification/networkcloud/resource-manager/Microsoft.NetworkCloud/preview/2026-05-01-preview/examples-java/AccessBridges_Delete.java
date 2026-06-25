
import com.azure.resourcemanager.networkcloud.models.AccessBridgeAllowedName;

/**
 * Samples for AccessBridges Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/AccessBridges_Delete.json
     */
    /**
     * Sample code: Delete access bridge.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void deleteAccessBridge(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.accessBridges().delete("resourceGroupName", AccessBridgeAllowedName.BASTION, null, null,
            com.azure.core.util.Context.NONE);
    }
}
