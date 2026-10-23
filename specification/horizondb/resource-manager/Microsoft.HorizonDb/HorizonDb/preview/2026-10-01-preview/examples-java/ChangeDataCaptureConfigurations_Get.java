
/**
 * Samples for ChangeDataCaptureConfigurations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ChangeDataCaptureConfigurations_Get.json
     */
    /**
     * Sample code: Get a Microsoft Fabric Mirroring configuration.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        getAMicrosoftFabricMirroringConfiguration(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.changeDataCaptureConfigurations().getWithResponse("example-resource-group", "example-cluster",
            "sales-database", com.azure.core.util.Context.NONE);
    }
}
