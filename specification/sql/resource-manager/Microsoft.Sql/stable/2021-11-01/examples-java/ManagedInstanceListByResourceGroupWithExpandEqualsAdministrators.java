
/**
 * Samples for ManagedInstances ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/
     * ManagedInstanceListByResourceGroupWithExpandEqualsAdministrators.json
     */
    /**
     * Sample code: List managed instances by resource group with $expand=administrators/activedirectory.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listManagedInstancesByResourceGroupWithExpandAdministratorsActivedirectory(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedInstances().listByResourceGroup("Test1", null,
            com.azure.core.util.Context.NONE);
    }
}
