/** Samples for KustoPoolAttachedDatabaseConfigurations ListByKustoPool. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsListByKustoPool.json
     */
    /**
     * Sample code: KustoPoolAttachedDatabaseConfigurationsListByKustoPool.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolAttachedDatabaseConfigurationsListByKustoPool(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolAttachedDatabaseConfigurations()
            .listByKustoPool("kustorptest", "kustoclusterrptest4", "kustorptest", com.azure.core.util.Context.NONE);
    }
}
