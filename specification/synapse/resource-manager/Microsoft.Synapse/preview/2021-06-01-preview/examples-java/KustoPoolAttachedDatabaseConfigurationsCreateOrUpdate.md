```java
import com.azure.resourcemanager.synapse.models.DefaultPrincipalsModificationKind;
import com.azure.resourcemanager.synapse.models.TableLevelSharingProperties;
import java.util.Arrays;

/** Samples for KustoPoolAttachedDatabaseConfigurations CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsCreateOrUpdate.json
     */
    /**
     * Sample code: KustoPoolAttachedDatabaseConfigurationsCreateOrUpdate.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolAttachedDatabaseConfigurationsCreateOrUpdate(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolAttachedDatabaseConfigurations()
            .define("attachedDatabaseConfigurations1")
            .withExistingKustoPool("kustorptest", "kustoclusterrptest4", "kustorptest")
            .withRegion("westus")
            .withDatabaseName("kustodatabase")
            .withKustoPoolResourceId(
                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/Workspaces/kustorptest/KustoPools/kustoclusterrptest4")
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

Read the [SDK documentation](https://github.com/Azure/azure-sdk-for-java/blob/azure-resourcemanager-synapse_1.0.0-beta.6/sdk/synapse/azure-resourcemanager-synapse/README.md) on how to add the SDK to your project and authenticate.
