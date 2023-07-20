/** Samples for AccessControlLists GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/managednetworkfabric/resource-manager/Microsoft.ManagedNetworkFabric/stable/2023-06-15/examples/AccessControlLists_Get_MaximumSet_Gen.json
     */
    /**
     * Sample code: AccessControlLists_Get_MaximumSet_Gen.
     *
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void accessControlListsGetMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager
            .accessControlLists()
            .getByResourceGroupWithResponse("example-rg", "example-acl", com.azure.core.util.Context.NONE);
    }
}
