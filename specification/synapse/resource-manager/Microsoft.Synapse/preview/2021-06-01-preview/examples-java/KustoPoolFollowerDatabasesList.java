/** Samples for KustoPools ListFollowerDatabases. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolFollowerDatabasesList.json
     */
    /**
     * Sample code: KustoPoolListFollowerDatabases.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolListFollowerDatabases(com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPools()
            .listFollowerDatabases(
                "kustorptest", "kustoclusterrptest4", "kustorptest", com.azure.core.util.Context.NONE);
    }
}
