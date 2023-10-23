import com.azure.resourcemanager.kusto.models.Language;

/** Samples for SandboxCustomImages CreateOrUpdate. */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2023-08-15/examples/KustoSandboxCustomImagesCreateOrUpdate.json
     */
    /**
     * Sample code: KustoSandboxCustomImagesCreateOrUpdate.
     *
     * @param manager Entry point to KustoManager.
     */
    public static void kustoSandboxCustomImagesCreateOrUpdate(com.azure.resourcemanager.kusto.KustoManager manager) {
        manager
            .sandboxCustomImages()
            .define("customImage8")
            .withExistingCluster("kustorptest", "kustoCluster")
            .withLanguage(Language.PYTHON)
            .withLanguageVersion("3.10.8")
            .withRequirementsFileContent("Requests")
            .create();
    }
}
