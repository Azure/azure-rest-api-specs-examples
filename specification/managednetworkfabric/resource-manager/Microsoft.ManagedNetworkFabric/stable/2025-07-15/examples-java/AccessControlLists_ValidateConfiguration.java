
/**
 * Samples for AccessControlLists ValidateConfiguration.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-07-15/AccessControlLists_ValidateConfiguration.json
     */
    /**
     * Sample code: AccessControlLists_ValidateConfiguration_MaximumSet_Gen.
     * 
     * @param manager Entry point to ManagedNetworkFabricManager.
     */
    public static void accessControlListsValidateConfigurationMaximumSetGen(
        com.azure.resourcemanager.managednetworkfabric.ManagedNetworkFabricManager manager) {
        manager.accessControlLists().validateConfiguration("example-rg", "example-acl",
            com.azure.core.util.Context.NONE);
    }
}
