
/**
 * Samples for NetAppResourceQuotaLimits List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/preview/2025-07-01-preview/examples/QuotaLimits_List.json
     */
    /**
     * Sample code: QuotaLimits.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void quotaLimits(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResourceQuotaLimits().list("eastus", com.azure.core.util.Context.NONE);
    }
}
