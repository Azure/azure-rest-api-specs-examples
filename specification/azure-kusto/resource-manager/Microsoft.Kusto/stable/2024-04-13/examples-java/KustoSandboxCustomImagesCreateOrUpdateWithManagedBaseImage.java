
import com.azure.resourcemanager.kusto.models.Language;

/**
 * Samples for SandboxCustomImages CreateOrUpdate.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/azure-kusto/resource-manager/Microsoft.Kusto/stable/2024-04-13/examples/
     * KustoSandboxCustomImagesCreateOrUpdateWithManagedBaseImage.json
     */
    /**
     * Sample code: KustoSandboxCustomImagesCreateOrUpdateWithManagedBaseImage.
     * 
     * @param manager Entry point to KustoManager.
     */
    public static void kustoSandboxCustomImagesCreateOrUpdateWithManagedBaseImage(
        com.azure.resourcemanager.kusto.KustoManager manager) {
        manager.sandboxCustomImages().define("customImage2").withExistingCluster("kustorptest", "kustoCluster")
            .withLanguage(Language.PYTHON).withBaseImageName("Python3_10_8").withRequirementsFileContent("Requests")
            .create();
    }
}
