import com.azure.core.util.Context;
import com.azure.resourcemanager.netapp.models.VolumeQuotaRule;

/** Samples for VolumeQuotaRules Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/netapp/resource-manager/Microsoft.NetApp/stable/2022-01-01/examples/VolumeQuotaRules_Update.json
     */
    /**
     * Sample code: VolumeQuotaRules_Update.
     *
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeQuotaRulesUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        VolumeQuotaRule resource =
            manager
                .volumeQuotaRules()
                .getWithResponse("myRG", "account-9957", "pool-5210", "volume-6387", "rule-0004", Context.NONE)
                .getValue();
        resource.update().withQuotaSizeInKiBs(100009L).apply();
    }
}
