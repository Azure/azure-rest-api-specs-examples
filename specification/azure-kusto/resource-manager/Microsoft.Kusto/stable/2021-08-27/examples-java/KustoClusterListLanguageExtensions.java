import com.azure.core.util.Context;

/** Samples for Clusters ListLanguageExtensions. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2021-08-27/examples/KustoClusterListLanguageExtensions.json
     */
    /**
     * Sample code: KustoClusterListLanguageExtensions.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterListLanguageExtensions(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().listLanguageExtensions("kustorptest", "kustoclusterrptest4", Context.NONE);
    }
}
