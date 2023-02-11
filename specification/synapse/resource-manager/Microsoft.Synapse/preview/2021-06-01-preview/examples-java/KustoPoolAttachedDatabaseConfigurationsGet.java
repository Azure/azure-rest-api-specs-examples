/** Samples for KustoPoolAttachedDatabaseConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/synapse/resource-manager/Microsoft.Synapse/preview/2021-06-01-preview/examples/KustoPoolAttachedDatabaseConfigurationsGet.json
     */
    /**
     * Sample code: KustoPoolAttachedDatabaseConfigurationsGet.
     *
     * @param manager Entry point to SynapseManager.
     */
    public static void kustoPoolAttachedDatabaseConfigurationsGet(
        com.azure.resourcemanager.synapse.SynapseManager manager) {
        manager
            .kustoPoolAttachedDatabaseConfigurations()
            .getWithResponse(
                "kustorptest",
                "kustoclusterrptest4",
                "attachedDatabaseConfigurations1",
                "kustorptest",
                com.azure.core.util.Context.NONE);
    }
}
