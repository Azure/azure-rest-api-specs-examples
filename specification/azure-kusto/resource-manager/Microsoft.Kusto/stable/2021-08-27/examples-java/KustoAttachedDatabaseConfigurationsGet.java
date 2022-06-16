import com.azure.core.util.Context;

/** Samples for AttachedDatabaseConfigurations Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoAttachedDatabaseConfigurationsGet.json
     */
    /**
     * Sample code: AttachedDatabaseConfigurationsGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void attachedDatabaseConfigurationsGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .attachedDatabaseConfigurations()
            .getWithResponse("kustorptest", "kustoclusterrptest4", "attachedDatabaseConfigurations1", Context.NONE);
    }
}
