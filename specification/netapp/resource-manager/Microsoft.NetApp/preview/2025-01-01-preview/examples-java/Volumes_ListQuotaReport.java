
/**
 * Samples for Volumes ListQuotaReport.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-01-01-preview/examples/
     * Volumes_ListQuotaReport.json
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
