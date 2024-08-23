
import com.azure.resourcemanager.machinelearning.models.PendingUploadRequestDto;
import com.azure.resourcemanager.machinelearning.models.PendingUploadType;

/**
 * Samples for RegistryModelVersions CreateOrGetStartPendingUpload.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/ModelVersion/createOrGetStartPendingUpload.json
     */
    /**
     * Sample code: CreateOrGetStartPendingUpload Registry Model Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrGetStartPendingUploadRegistryModelVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryModelVersions()
            .createOrGetStartPendingUploadWithResponse(
                "test-rg", "registryName", "string", "string", new PendingUploadRequestDto()
                    .withPendingUploadId("string").withPendingUploadType(PendingUploadType.TEMPORARY_BLOB_REFERENCE),
                com.azure.core.util.Context.NONE);
    }
}
