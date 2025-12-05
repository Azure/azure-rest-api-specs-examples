
import com.azure.resourcemanager.netapp.models.VolumeQuotaRule;

/**
 * Samples for VolumeQuotaRules Update.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/VolumeQuotaRules_Update.json
     */
    /**
     * Sample code: VolumeQuotaRules_Update.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void volumeQuotaRulesUpdate(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        VolumeQuotaRule resource = manager.volumeQuotaRules().getWithResponse("myRG", "account-9957", "pool-5210",
            "volume-6387", "rule-0004", com.azure.core.util.Context.NONE).getValue();
        resource.update().withQuotaSizeInKiBs(100009L).apply();
    }
}
