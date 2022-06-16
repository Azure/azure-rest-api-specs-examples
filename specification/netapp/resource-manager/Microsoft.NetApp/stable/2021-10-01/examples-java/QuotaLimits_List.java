import com.azure.core.util.Context;

/** Samples for NetAppResourceQuotaLimits List. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2021-10-01/examples/QuotaLimits_List.json
     */
    /**
     * Sample code: QuotaLimits.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void quotaLimits(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResourceQuotaLimits().list("eastus", Context.NONE);
    }
}
