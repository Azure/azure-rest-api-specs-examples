
/**
 * Samples for Volumes ListQuotaReport.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/Volumes_ListQuotaReport.json
     */
    /**
     * Sample code: ListQuotaReport.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void listQuotaReport(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumes().listQuotaReport("myRG", "account1", "pool1", "volume1", com.azure.core.util.Context.NONE);
    }
}
