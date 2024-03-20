
import com.azure.resourcemanager.iotfirmwaredefense.models.GenerateUploadUrlRequest;

/**
 * Samples for Workspaces GenerateUploadUrl.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/stable/2024-01-10/examples/
     * Workspaces_GenerateUploadUrl_MaximumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_GenerateUploadUrl_MaximumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void workspacesGenerateUploadUrlMaximumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.workspaces().generateUploadUrlWithResponse("rgworkspaces", "E___-3",
            new GenerateUploadUrlRequest().withFirmwareId("ytsfprbywi"), com.azure.core.util.Context.NONE);
    }
}
