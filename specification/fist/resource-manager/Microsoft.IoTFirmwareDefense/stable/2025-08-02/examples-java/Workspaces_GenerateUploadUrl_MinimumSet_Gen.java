
import com.azure.resourcemanager.iotfirmwaredefense.models.GenerateUploadUrlRequest;

/**
 * Samples for Workspaces GenerateUploadUrl.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-08-02/Workspaces_GenerateUploadUrl_MinimumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_GenerateUploadUrl_MinimumSet_Gen.
     * 
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void workspacesGenerateUploadUrlMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager.workspaces().generateUploadUrlWithResponse("rgworkspaces", "default",
            new GenerateUploadUrlRequest().withFirmwareId("00000000-0000-0000-0000-000000000000"),
            com.azure.core.util.Context.NONE);
    }
}
