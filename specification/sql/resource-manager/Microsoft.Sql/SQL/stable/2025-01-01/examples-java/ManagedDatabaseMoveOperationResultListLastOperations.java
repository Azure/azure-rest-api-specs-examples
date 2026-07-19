
/**
 * Samples for ManagedDatabaseMoveOperations ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseMoveOperationResultListLastOperations.json
     */
    /**
     * Sample code: Gets the latest managed database move operations for each database under specified subscription,
     * resource group and location.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsTheLatestManagedDatabaseMoveOperationsForEachDatabaseUnderSpecifiedSubscriptionResourceGroupAndLocation(
            com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseMoveOperations().listByLocation("rg1", "westeurope", null, null,
            com.azure.core.util.Context.NONE);
    }
}
