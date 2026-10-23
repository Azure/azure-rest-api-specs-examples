
/**
 * Samples for ChangeDataCaptureCapabilities Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ChangeDataCaptureCapabilities_Delete.json
     */
    /**
     * Sample code: Delete the cluster change data capture capability.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        deleteTheClusterChangeDataCaptureCapability(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.changeDataCaptureCapabilities().delete("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
