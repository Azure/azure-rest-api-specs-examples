import com.azure.core.util.Context;

/** Samples for KustoPoolAttachedDatabaseConfigurations Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsDelete.json
     */
    /**
     * Sample code: KustoPoolAttachedDatabaseConfigurationsDelete.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolAttachedDatabaseConfigurationsDelete(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolAttachedDatabaseConfigurations()
            .delete(
                "kustorptest", "kustoclusterrptest4", "attachedDatabaseConfigurations1", "kustorptest", Context.NONE);
    }
}
