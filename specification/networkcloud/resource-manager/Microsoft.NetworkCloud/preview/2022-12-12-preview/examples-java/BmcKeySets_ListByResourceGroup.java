/** Samples for BmcKeySets ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/networkcloud/resource-manager/Microsoft.NetworkCloud/preview/2022-12-12-preview/examples/BmcKeySets_ListByResourceGroup.json
     */
    /**
     * Sample code: List baseboard management controller key set of cluster for resource group.
     *
     * @param manager Entry point to NetworkCloudManager.
     */
    public static void listBaseboardManagementControllerKeySetOfClusterForResourceGroup(
        com.azure.resourcemanager.networkcloud.NetworkCloudManager manager) {
        manager.bmcKeySets().listByResourceGroup("resourceGroupName", "clusterName", com.azure.core.util.Context.NONE);
    }
}
