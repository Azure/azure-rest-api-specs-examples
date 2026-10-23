
/**
 * Samples for ChangeDataCaptureConfigurations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ChangeDataCaptureConfigurations_List.json
     */
    /**
     * Sample code: List Microsoft Fabric Mirroring configurations.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listMicrosoftFabricMirroringConfigurations(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.changeDataCaptureConfigurations().list("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
