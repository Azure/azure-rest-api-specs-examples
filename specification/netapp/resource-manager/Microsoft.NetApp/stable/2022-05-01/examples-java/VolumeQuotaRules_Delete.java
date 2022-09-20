import com.azure.core.util.Context;

/** Samples for VolumeQuotaRules Delete. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-05-01/examples/VolumeQuotaRules_Delete.json
     */
    /**
     * Sample code: VolumeQuotaRules_Delete.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeQuotaRulesDelete(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager
            .volumeQuotaRules()
            .delete("myRG", "account-9957", "pool-5210", "volume-6387", "rule-0004", Context.NONE);
    }
}
