
/**
 * Samples for AccessControlLists Delete.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/
     * AccessControlLists_Delete_MaximumSet_Gen.json
     */
    /**
     * Sample code: AccessControlLists_Delete_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void accessControlListsDeleteMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.accessControlLists().delete("example-rg", "example-acl", com.azure.core.util.Context.NONE);
    }
}
