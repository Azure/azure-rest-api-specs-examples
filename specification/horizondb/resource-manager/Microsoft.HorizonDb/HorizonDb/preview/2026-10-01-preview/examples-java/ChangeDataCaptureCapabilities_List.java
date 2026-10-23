
/**
 * Samples for ChangeDataCaptureCapabilities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ChangeDataCaptureCapabilities_List.json
     */
    /**
     * Sample code: List the optional cluster change data capture capability.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void listTheOptionalClusterChangeDataCaptureCapability(
        com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.changeDataCaptureCapabilities().list("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
