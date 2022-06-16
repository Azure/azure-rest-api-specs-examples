import com.azure.core.util.Context;

/** Samples for Clusters ListLanguageExtensions. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2022-02-01/examples/KustoClusterListLanguageExtensions.json
     */
    /**
     * Sample code: KustoClusterListLanguageExtensions.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoClusterListLanguageExtensions(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.clusters().listLanguageExtensions("kustorptest", "kustoCluster", Context.NONE);
    }
}
