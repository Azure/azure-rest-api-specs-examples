import com.azure.resourcemanager.kusto.models.DefaultPrincipalsModificationKind;
import com.azure.resourcemanager.kusto.models.TableLevelSharingProperties;
import java.util.Arrays;

/** Samples for AttachedDatabaseConfigurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
     */
    /**
     * Sample code: AttachedDatabaseConfigurationsCreateOrUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void attachedDatabaseConfigurationsCreateOrUpdate(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .attachedDatabaseConfigurations()
            .define("attachedDatabaseConfigurations1")
            .withExistingCluster("kustorptest", "kustoclusterrptest4")
            .withRegion("westus")
            .withDatabaseName("kustodatabase")
            .withClusterResourceId(
                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/KustoClusterLeader")
            .withDefaultPrincipalsModificationKind(DefaultPrincipalsModificationKind.UNION)
            .withTableLevelSharingProperties(
                new TableLevelSharingProperties()
                    .withTablesToInclude(Arrays.asList("Table1"))
                    .withTablesToExclude(Arrays.asList("Table2"))
                    .withExternalTablesToInclude(Arrays.asList("ExternalTable1"))
                    .withExternalTablesToExclude(Arrays.asList("ExternalTable2"))
                    .withMaterializedViewsToInclude(Arrays.asList("MaterializedViewTable1"))
                    .withMaterializedViewsToExclude(Arrays.asList("MaterializedViewTable2")))
            .create();
    }
}
