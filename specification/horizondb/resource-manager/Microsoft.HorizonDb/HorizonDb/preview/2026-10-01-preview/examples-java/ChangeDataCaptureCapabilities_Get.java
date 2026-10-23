
/**
 * Samples for ChangeDataCaptureCapabilities Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-10-01-preview/ChangeDataCaptureCapabilities_Get.json
     */
    /**
     * Sample code: Get the cluster change data capture capability.
     * 
     * @param manager Entry point to HorizonDbManager.
     */
    public static void
        getTheClusterChangeDataCaptureCapability(com.azure.resourcemanager.horizondb.HorizonDbManager manager) {
        manager.changeDataCaptureCapabilities().getWithResponse("example-resource-group", "example-cluster",
            com.azure.core.util.Context.NONE);
    }
}
