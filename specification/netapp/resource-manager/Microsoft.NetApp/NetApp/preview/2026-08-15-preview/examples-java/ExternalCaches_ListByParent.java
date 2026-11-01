
/**
 * Samples for ExternalCaches ListByParent.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-08-15-preview/ExternalCaches_ListByParent.json
     */
    /**
     * Sample code: ExternalCaches_ListByParent.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void externalCachesListByParent(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.externalCaches().listByParent("contosoNetAppRg", "contosoExternalCluster", "contosoExternalSvm",
            com.azure.core.util.Context.NONE);
    }
}
