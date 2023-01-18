import com.azure.resourcemanager.devtestlabs.models.FileUploadOptions;
import com.azure.resourcemanager.devtestlabs.models.GenerateArmTemplateRequest;

/** Samples for Artifacts GenerateArmTemplate. */
public final class Main {
    /*
     * x-ms-original-file: specification/devtestlabs/resource-manager/Microsoft.DevTestLab/stable/2018-09-15/examples/Artifacts_GenerateArmTemplate.json
     */
    /**
     * Sample code: Artifacts_GenerateArmTemplate.
     *
     * @param manager Entry point to DevTestLabsManager.
     */
    public static void artifactsGenerateArmTemplate(com.azure.resourcemanager.devtestlabs.DevTestLabsManager manager) {
        manager
            .artifacts()
            .generateArmTemplateWithResponse(
                "resourceGroupName",
                "{labName}",
                "{artifactSourceName}",
                "{artifactName}",
                new GenerateArmTemplateRequest()
                    .withVirtualMachineName("{vmName}")
                    .withLocation("{location}")
                    .withFileUploadOptions(FileUploadOptions.NONE),
                com.azure.core.util.Context.NONE);
    }
}
