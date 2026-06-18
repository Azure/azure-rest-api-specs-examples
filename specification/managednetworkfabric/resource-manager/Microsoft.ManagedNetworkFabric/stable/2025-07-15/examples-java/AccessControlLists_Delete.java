
/**
 * Samples for AccessControlLists Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/AccessControlLists_Delete.json
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
