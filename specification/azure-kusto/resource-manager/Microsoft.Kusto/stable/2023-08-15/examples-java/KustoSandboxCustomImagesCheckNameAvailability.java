import com.azure.resourcemanager.kusto.models.SandboxCustomImagesCheckNameRequest;

/** Samples for SandboxCustomImages CheckNameAvailability. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoSandboxCustomImagesCheckNameAvailability.json
     */
    /**
     * Sample code: KustoSandboxCustomImagesCheckNameAvailability.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoSandboxCustomImagesCheckNameAvailability(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .sandboxCustomImages()
            .checkNameAvailabilityWithResponse(
                "kustorptest",
                "kustoCluster",
                new SandboxCustomImagesCheckNameRequest().withName("sandboxCustomImage1"),
                com.azure.core.util.Context.NONE);
    }
}
