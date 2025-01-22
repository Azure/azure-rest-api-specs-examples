
/**
 * Samples for QuotaRequestStatus Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2024-12-18-preview/examples/
     * getQuotaRequestStatusById.json
     */
    /**
     * Sample code: QuotaRequestStatus.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void quotaRequestStatus(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.quotaRequestStatus().getWithResponse("2B5C8515-37D8-4B6A-879B-CD641A2CF605",
            "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus",
            com.azure.core.util.Context.NONE);
    }
}
