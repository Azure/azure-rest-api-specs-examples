
/**
 * Samples for ChangeDataCaptureDestinations Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ChangeDataCaptureDestinations_Delete.json
     */
    /**
     * Sample code: Delete a Microsoft Fabric change data capture destination.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void deleteAMicrosoftFabricChangeDataCaptureDestination(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.changeDataCaptureDestinations().delete("example-resource-group", "example-cluster", "fabric-mirroring",
            com.azure.core.util.Context.NONE);
    }
}
