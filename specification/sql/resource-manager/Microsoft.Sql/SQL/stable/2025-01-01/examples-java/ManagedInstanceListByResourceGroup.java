
/**
 * Samples for ManagedInstances ListByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceListByResourceGroup.json
     */
    /**
     * Sample code: List managed instances by resource group.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void listManagedInstancesByResourceGroup(com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().listByResourceGroup("Test1", null,
            com.azure.core.util.Context.NONE);
    }
}
