
/**
 * Samples for ExternalSvms Get.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalSvms_Get.json
     */
    /**
     * Sample code: ExternalSvms_Get.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void externalSvmsGet(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalSvms().getWithResponse("contosoNetAppRg", "contoso-external-cluster", "contoso-svm",
            com.azure.core.util.Context.NONE);
    }
}
