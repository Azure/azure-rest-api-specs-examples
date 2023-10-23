/** Samples for SandboxCustomImages Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoSandboxCustomImagesGet.json
     */
    /**
     * Sample code: KustoSandboxCustomImagesGet.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoSandboxCustomImagesGet(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .sandboxCustomImages()
            .getWithResponse("kustorptest", "kustoCluster", "customImage8", com.azure.core.util.Context.NONE);
    }
}
