/** Samples for Racks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2023-05-01-preview/examples/Racks_ListByResourceGroup.json
     */
    /**
     * Sample code: List racks for resource group.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listRacksForResourceGroup(com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.racks().listByResourceGroup("resourceGroupName", com.azure.core.util.Context.NONE);
    }
}
