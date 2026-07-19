
/**
 * Samples for ManagedDatabaseMoveOperations ListByLocation.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseMoveOperationResultList.json
     */
    /**
     * Sample code: Gets all managed database move operations for specified subscription, resource group and location.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void getsAllManagedDatabaseMoveOperationsForSpecifiedSubscriptionResourceGroupAndLocation(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabaseMoveOperations().listByLocation("rg1", "westeurope", null, null,
            com.azure.core.util.Context.NONE);
    }
}
