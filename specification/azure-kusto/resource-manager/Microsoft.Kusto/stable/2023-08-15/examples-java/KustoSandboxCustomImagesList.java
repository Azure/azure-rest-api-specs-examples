/** Samples for SandboxCustomImages ListByCluster. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoSandboxCustomImagesList.json
     */
    /**
     * Sample code: KustoSandboxCustomImagesListByCluster.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoSandboxCustomImagesListByCluster(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.sandboxCustomImages().listByCluster("kustorptest", "kustoCluster", com.azure.core.util.Context.NONE);
    }
}
