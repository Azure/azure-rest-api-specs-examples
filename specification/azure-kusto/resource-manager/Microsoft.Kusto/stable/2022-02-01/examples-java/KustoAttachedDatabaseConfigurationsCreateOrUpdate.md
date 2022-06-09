```java
import com.azure.resourcemanager.kusto.models.DefaultPrincipalsModificationKind;
import com.azure.resourcemanager.kusto.models.TableLevelSharingProperties;
import java.util.Arrays;

/** Samples for AttachedDatabaseConfigurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoAttachedDatabaseConfigurationsCreateOrUpdate.json
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
            .define("attachedDatabaseConfigurationsTest")
            .withExistingCluster("kustorptest", "kustoCluster2")
            .withRegion("westus")
            .withDatabaseName("kustodatabase")
            .withClusterResourceId(
                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Kusto/Clusters/kustoCluster2")
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
```

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-kusto_1.0.0-beta.4/sdk/kusto/azure-resourcemanager-kusto/README.md) on how to add the SDK to your project and authenticate.
