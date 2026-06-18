
/**
 * Samples for AccessControlLists ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/AccessControlLists_ListByResourceGroup.json
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
