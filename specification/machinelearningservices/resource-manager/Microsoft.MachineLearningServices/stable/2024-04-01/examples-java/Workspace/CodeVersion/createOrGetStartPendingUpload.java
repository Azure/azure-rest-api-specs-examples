
import com.azure.resourcemanager.machinelearning.models.PendingUploadRequestDto;
import com.azure.resourcemanager.machinelearning.models.PendingUploadType;

/**
 * Samples for CodeVersions CreateOrGetStartPendingUpload.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/machinelearningservices/resource-manager/Microsoft.MachineLearningServices/stable/2024-04-01/
     * examples/Workspace/CodeVersion/createOrGetStartPendingUpload.json
     */
    /**
     * Sample code: CreateOrGetStartPendingUpload Workspace Code Version.
     * 
     * @param manager Entry point to MachineLearningManager.
     */
    public static void createOrGetStartPendingUploadWorkspaceCodeVersion(
        com.azure.resourcemanager.machinelearning.MachineLearningManager manager) {
        manager.codeVersions().createOrGetStartPendingUploadWithResponse("test-rg", "my-aml-workspace", "string",
            "string",
            new PendingUploadRequestDto().withPendingUploadId("string").withPendingUploadType(PendingUploadType.NONE),
            com.azure.core.util.Context.NONE);
    }
}
