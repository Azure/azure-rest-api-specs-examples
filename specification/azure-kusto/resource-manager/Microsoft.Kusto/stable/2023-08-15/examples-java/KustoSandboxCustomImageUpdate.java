import com.azure.resourcemanager.kusto.models.Language;
import com.azure.resourcemanager.kusto.models.SandboxCustomImage;

/** Samples for SandboxCustomImages Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoSandboxCustomImageUpdate.json
     */
    /**
     * Sample code: KustoSandboxCustomImagesUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoSandboxCustomImagesUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        SandboxCustomImage resource =
            manager
                .sandboxCustomImages()
                .getWithResponse("kustorptest", "kustoCluster", "customImage8", com.azure.core.util.Context.NONE)
                .getValue();
        resource
            .update()
            .withLanguage(Language.PYTHON)
            .withLanguageVersion("3.10.8")
            .withRequirementsFileContent("Requests")
            .apply();
    }
}
