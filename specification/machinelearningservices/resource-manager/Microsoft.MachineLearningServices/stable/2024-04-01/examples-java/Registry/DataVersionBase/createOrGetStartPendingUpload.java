
import com.azure.resourcemanager.machinelearning.models.PendingUploadRequestDto;
import com.azure.resourcemanager.machinelearning.models.PendingUploadType;

/**
 * Samples for RegistryDataVersions CreateOrGetStartPendingUpload.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/DataVersionBase/createOrGetStartPendingUpload.json
     */
    /**
     * Sample code: CreateOrGetStartPendingUpload Registry Data Version Base.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrGetStartPendingUploadRegistryDataVersionBase(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryDataVersions().createOrGetStartPendingUploadWithResponse("test-rg", "registryName", "string",
            "string",
            new PendingUploadRequestDto().withPendingUploadId("string").withPendingUploadType(PendingUploadType.NONE),
            com.azure.core.util.Context.NONE);
    }
}
