
import com.azure.resourcemanager.kusto.models.DatabaseInviteFollowerRequest;
import com.azure.resourcemanager.kusto.models.TableLevelSharingProperties;
import java.util.Arrays;

/**
 * Samples for DatabaseOperation InviteFollower.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/KustoDatabaseInviteFollower
     * .json
     */
    /**
     * Sample code: KustoDatabaseInviteFollower.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoDatabaseInviteFollower(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.databaseOperations().inviteFollowerWithResponse("kustorptest", "kustoCluster", "database",
            new DatabaseInviteFollowerRequest().withInviteeEmail("invitee@contoso.com")
                .withTableLevelSharingProperties(new TableLevelSharingProperties()
                    .withTablesToInclude(Arrays.asList("Table1")).withTablesToExclude(Arrays.asList("Table2"))
                    .withExternalTablesToInclude(Arrays.asList("ExternalTable*"))
                    .withExternalTablesToExclude(Arrays.asList())
                    .withMaterializedViewsToInclude(Arrays.asList("MaterializedViewTable1"))
                    .withMaterializedViewsToExclude(Arrays.asList("MaterializedViewTable2"))
                    .withFunctionsToInclude(Arrays.asList("functionsToInclude1"))
                    .withFunctionsToExclude(Arrays.asList("functionsToExclude2"))),
            com.azure.core.util.Context.NONE);
    }
}
