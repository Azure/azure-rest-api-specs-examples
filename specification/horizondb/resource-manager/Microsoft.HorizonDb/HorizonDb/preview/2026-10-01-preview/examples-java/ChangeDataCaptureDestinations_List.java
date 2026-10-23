
/**
 * Samples for ChangeDataCaptureDestinations List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ChangeDataCaptureDestinations_List.json
     */
    /**
     * Sample code: List Microsoft Fabric change data capture destinations.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        listMicrosoftFabricChangeDataCaptureDestinations(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.changeDataCaptureDestinations().list("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
