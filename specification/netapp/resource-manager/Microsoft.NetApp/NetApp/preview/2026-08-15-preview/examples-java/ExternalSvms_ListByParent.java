
/**
 * Samples for ExternalSvms ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalSvms_ListByParent.json
     */
    /**
     * Sample code: ExternalSvms_ListByParent.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void externalSvmsListByParent(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalSvms().listByParent("contosoNetAppRg", "contoso-external-cluster",
            com.azure.core.util.Context.NONE);
    }
}
