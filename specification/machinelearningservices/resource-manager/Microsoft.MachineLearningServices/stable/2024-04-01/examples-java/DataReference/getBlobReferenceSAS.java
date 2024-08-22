
import com.azure.resourcemanager.machinelearning.models.GetBlobReferenceSasRequestDto;

/**
 * Samples for RegistryDataReferences GetBlobReferenceSas.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/DataReference/getBlobReferenceSAS.json
     */
    /**
     * Sample code: GetBlobReferenceSAS Data Reference.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void
        getBlobReferenceSASDataReference(com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryDataReferences().getBlobReferenceSasWithResponse("test-rg", "registryName", "string", "string",
            new GetBlobReferenceSasRequestDto().withAssetId("string").withBlobUri("https://www.contoso.com/example"),
            com.azure.core.util.Context.NONE);
    }
}
