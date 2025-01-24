
/**
 * Samples for SandboxCustomImages Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/
     * KustoSandboxCustomImageDelete.json
     */
    /**
     * Sample code: SandboxCustomImagesDelete.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void sandboxCustomImagesDelete(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.sandboxCustomImages().delete("kustorptest", "kustoCluster", "customImage8",
            com.azure.core.util.Context.NONE);
    }
}
