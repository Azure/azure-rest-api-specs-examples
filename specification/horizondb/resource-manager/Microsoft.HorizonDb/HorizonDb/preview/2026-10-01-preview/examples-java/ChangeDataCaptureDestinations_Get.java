
/**
 * Samples for ChangeDataCaptureDestinations Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ChangeDataCaptureDestinations_Get.json
     */
    /**
     * Sample code: Get a Microsoft Fabric change data capture destination.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        getAMicrosoftFabricChangeDataCaptureDestination(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.changeDataCaptureDestinations().getWithResponse("example-resource-group", "example-cluster",
            "fabric-mirroring", com.azure.core.util.Context.NONE);
    }
}
