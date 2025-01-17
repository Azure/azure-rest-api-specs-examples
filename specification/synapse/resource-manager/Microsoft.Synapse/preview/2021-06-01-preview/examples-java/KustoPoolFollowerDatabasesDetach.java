
import com.azure.resourcemanager.synapse.fluent.models.FollowerDatabaseDefinitionInner;

/**
 * Samples for KustoPools DetachFollowerDatabases.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/
     * KustoPoolFollowerDatabasesDetach.json
     */
    /**
     * Sample code: KustoPoolDetachFollowerDatabases.
     * 
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolDetachFollowerDatabases(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager.kustoPools().detachFollowerDatabases("kustorptest", "kustoclusterrptest4", "kustorptest",
            new FollowerDatabaseDefinitionInner().withKustoPoolResourceId(
                "/subscriptions/12345678-1234-1234-1234-123456789098/resourceGroups/kustorptest/providers/Microsoft.Synapse/workspaces/kustorptest/kustoPools/leader4")
                .withAttachedDatabaseConfigurationName("myAttachedDatabaseConfiguration"),
            com.azure.core.util.Context.NONE);
    }
}
