
/**
 * Samples for ExternalSvms Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalSvms_Delete.json
     */
    /**
     * Sample code: ExternalSvms_Delete.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void externalSvmsDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalSvms().delete("contosoNetAppRg", "contoso-external-cluster", "contoso-svm",
            com.azure.core.util.Context.NONE);
    }
}
