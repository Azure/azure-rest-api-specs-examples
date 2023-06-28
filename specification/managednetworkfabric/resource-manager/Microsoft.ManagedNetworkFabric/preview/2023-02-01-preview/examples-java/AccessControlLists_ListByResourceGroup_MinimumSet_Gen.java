/** Samples for AccessControlLists ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/preview/2023-02-01-preview/examples/AccessControlLists_ListByResourceGroup_MinimumSet_Gen.json
     */
    /**
     * Sample code: AccessControlLists_ListByResourceGroup_MinimumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void accessControlListsListByResourceGroupMinimumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.accessControlLists().listByResourceGroup("resourceGroupName", com.azure.core.util.Context.NONE);
    }
}
