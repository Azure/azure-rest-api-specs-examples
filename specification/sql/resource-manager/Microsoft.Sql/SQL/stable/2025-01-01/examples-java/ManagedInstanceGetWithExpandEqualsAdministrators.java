
/**
 * Samples for ManagedInstances GetByResourceGroup.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedInstanceGetWithExpandEqualsAdministrators.json
     */
    /**
     * Sample code: Get managed instance with $expand=administrators/activedirectory.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getManagedInstanceWithExpandAdministratorsActivedirectory(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedInstances().getByResourceGroupWithResponse("testrg", "testinstance", null,
            com.azure.core.util.Context.NONE);
    }
}
