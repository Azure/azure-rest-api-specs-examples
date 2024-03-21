
/**
 * Samples for VolumeQuotaRules ListByVolume.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/netapp/resource-manager/Microsoft.NetApp/stable/2023-07-01/examples/VolumeQuotaRules_List.json
     */
    /**
     * Sample code: VolumeQuotaRules_List.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeQuotaRulesList(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.volumeQuotaRules().listByVolume("myRG", "account-9957", "pool-5210", "volume-6387",
            com.azure.core.util.Context.NONE);
    }
}
