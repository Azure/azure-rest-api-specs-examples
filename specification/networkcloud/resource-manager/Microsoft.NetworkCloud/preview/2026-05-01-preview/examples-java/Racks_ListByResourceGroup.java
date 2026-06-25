
/**
 * Samples for Racks ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-05-01-preview/Racks_ListByResourceGroup.json
     */
    /**
     * Sample code: List racks for resource group.
     * 
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listRacksForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.racks().listByResourceGroup("resourceGroupName", null, null, com.azure.core.util.Context.NONE);
    }
}
