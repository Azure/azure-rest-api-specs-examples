
import com.azure.resourcemanager.networkcloud.models.AccessBridgeAllowedName;

/**
 * Samples for AccessBridges GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/AccessBridges_Get.json
     */
    /**
     * Sample code: Get access bridge.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void getAccessBridge(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.accessBridges().getByResourceGroupWithResponse("resourceGroupName", AccessBridgeAllowedName.BASTION,
            com.azure.core.util.Context.NONE);
    }
}
