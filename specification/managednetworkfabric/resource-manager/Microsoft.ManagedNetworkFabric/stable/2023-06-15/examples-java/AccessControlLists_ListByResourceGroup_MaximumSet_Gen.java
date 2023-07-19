/** Samples for AccessControlLists ListByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/AccessControlLists_ListByResourceGroup_MaximumSet_Gen.json
     */
    /**
     * Sample code: AccessControlLists_ListByResourceGroup_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void accessControlListsListByResourceGroupMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.accessControlLists().listByResourceGroup("example-rg", com.azure.core.util.Context.NONE);
    }
}
