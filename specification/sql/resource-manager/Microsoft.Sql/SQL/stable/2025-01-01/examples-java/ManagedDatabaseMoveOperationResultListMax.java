
/**
 * Samples for ManagedDatabaseMoveOperations ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseMoveOperationResultListMax.json
     */
    /**
     * Sample code: Gets the latest managed database move operations for each database under specified subscription,
     * resource group and location, filtered by operation type.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void
        getsTheLatestManagedDatabaseMoveOperationsForEachDatabaseUnderSpecifiedSubscriptionResourceGroupAndLocationFilteredByOperationType(
            com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseMoveOperations().listByLocation("rg1", "westeurope", null,
            "Properties/Operation eq 'StartManagedInstanceDatabaseMove'", com.azure.core.util.Context.NONE);
    }
}
