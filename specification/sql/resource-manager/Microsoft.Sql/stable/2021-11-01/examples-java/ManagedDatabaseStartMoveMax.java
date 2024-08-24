
import com.azure.resourcemanager.sql.models.ManagedDatabaseStartMoveDefinition;
import com.azure.resourcemanager.sql.models.MoveOperationMode;

/**
 * Samples for ManagedDatabases StartMove.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseStartMoveMax.json
     */
    /**
     * Sample code: Starts a managed database move with all optional parameters specified.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void startsAManagedDatabaseMoveWithAllOptionalParametersSpecified(
        com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabases().startMove("group1", "testInstanceSrc",
            "testDatabase",
            new ManagedDatabaseStartMoveDefinition().withDestinationManagedDatabaseId(
                "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/managedInstances/testInstanceTgt/databases/testDatabase")
                .withOperationMode(MoveOperationMode.COPY),
            com.azure.core.util.Context.NONE);
    }
}
