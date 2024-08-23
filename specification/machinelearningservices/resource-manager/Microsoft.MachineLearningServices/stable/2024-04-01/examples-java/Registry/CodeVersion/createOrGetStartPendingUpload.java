
import com.azure.resourcemanager.machinelearning.models.PendingUploadRequestDto;
import com.azure.resourcemanager.machinelearning.models.PendingUploadType;

/**
 * Samples for RegistryCodeVersions CreateOrGetStartPendingUpload.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Registry/CodeVersion/createOrGetStartPendingUpload.json
     */
    /**
     * Sample code: CreateOrGetStartPendingUpload Registry Code Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrGetStartPendingUploadRegistryCodeVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.registryCodeVersions()
            .createOrGetStartPendingUploadWithResponse(
                "test-rg", "registryName", "string", "string", new PendingUploadRequestDto()
                    .withPendingUploadId("string").withPendingUploadType(PendingUploadType.TEMPORARY_BLOB_REFERENCE),
                com.azure.core.util.Context.NONE);
    }
}
