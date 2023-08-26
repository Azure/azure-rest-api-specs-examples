/** Samples for Racks ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/stable/2023-07-01/examples/Racks_ListByResourceGroup.json
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
