
import com.azure.resourcemanager.netapp.models.CheckQuotaNameResourceTypes;
import com.azure.resourcemanager.netapp.models.QuotaAvailabilityRequest;

/**
 * Samples for NetAppResource CheckQuotaAvailability.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-09-01-preview/CheckQuotaAvailability.json
     */
    /**
     * Sample code: CheckQuotaAvailability.
     * 
     * @param manager Entry point to NetAppFilesManager.
     */
    public static void checkQuotaAvailability(com.azure.resourcemanager.netapp.NetAppFilesManager manager) {
        manager.netAppResources().checkQuotaAvailabilityWithResponse("eastus",
            new QuotaAvailabilityRequest().withName("resource1")
                .withType(CheckQuotaNameResourceTypes.MICROSOFT_NET_APP_NET_APP_ACCOUNTS).withResourceGroup("myRG"),
            com.azure.core.util.Context.NONE);
    }
}
