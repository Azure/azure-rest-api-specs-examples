
import com.azure.resourcemanager.sql.models.ManagedDatabaseMoveDefinition;

/**
 * Samples for ManagedDatabases CompleteMove.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/sql/resource-manager/Microsoft.Sql/stable/2021-11-01/examples/ManagedDatabaseCompleteMove.json
     */
    /**
     * Sample code: Completes a managed database move.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void completesAManagedDatabaseMove(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.sqlServers().manager().serviceClient().getManagedDatabases().completeMove("group1", "testInstanceSrc",
            "testDatabase",
            new ManagedDatabaseMoveDefinition().withDestinationManagedDatabaseId(
                "subscriptions/00000000-1111-2222-3333-444444444444/resourceGroups/group1/providers/Microsoft.Sql/managedInstances/testInstanceTgt/databases/testDatabase"),
            com.azure.core.util.Context.NONE);
    }
}
