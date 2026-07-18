
import com.azure.resourcemanager.sql.models.ManagedDatabaseStartMoveDefinition;
import com.azure.resourcemanager.sql.models.MoveOperationMode;

/**
 * Samples for ManagedDatabases StartMove.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-01-01/ManagedDatabaseStartMoveMax.json
     */
    /**
     * Sample code: Starts a managed database move with all optional parameters specified.
     * 
     * @param manager Entry point to SqlServerManager.
     */
    public static void startsAManagedDatabaseMoveWithAllOptionalParametersSpecified(
        com.azure.resourcemanager.sql.SqlServerManager manager) {
        manager.serviceClient().getManagedDatabases().startMove("group1", "testInstanceSrc", "testDatabase",
            new ManagedDatabaseStartMoveDefinition().withDestinationManagedDatabaseId(
                "/subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/managedInstances/testInstanceTgt/databases/testDatabase")
                .withOperationMode(MoveOperationMode.COPY),
            com.azure.core.util.Context.NONE);
    }
}
