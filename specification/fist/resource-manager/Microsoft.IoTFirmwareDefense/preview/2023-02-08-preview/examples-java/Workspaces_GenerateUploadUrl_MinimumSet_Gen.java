import com.azure.resourcemanager.iotfirmwaredefense.models.GenerateUploadUrlRequest;

/** Samples for Workspaces GenerateUploadUrl. */
public final class Main {
    /*
     * x-ms-original-file: specification/fist/resource-manager/Microsoft.IoTFirmwareDefense/preview/2023-02-08-preview/examples/Workspaces_GenerateUploadUrl_MinimumSet_Gen.json
     */
    /**
     * Sample code: Workspaces_GenerateUploadUrl_MinimumSet_Gen.
     *
     * @param manager Entry point to IoTFirmwareDefenseManager.
     */
    public static void workspacesGenerateUploadUrlMinimumSetGen(
        com.azure.resourcemanager.iotfirmwaredefense.IoTFirmwareDefenseManager manager) {
        manager
            .workspaces()
            .generateUploadUrlWithResponse(
                "rgworkspaces", "E___-3", new GenerateUploadUrlRequest(), com.azure.core.util.Context.NONE);
    }
}
