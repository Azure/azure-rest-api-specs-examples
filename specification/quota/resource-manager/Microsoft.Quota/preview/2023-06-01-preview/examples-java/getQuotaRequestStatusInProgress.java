
/**
 * Samples for QuotaRequestStatus Get.
 */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/preview/2023-06-01-preview/examples/
     * getQuotaRequestStatusInProgress.json
     */
    /**
     * Sample code: QuotaRequestInProgress.
     * 
     * @param manager Entry point to QuotaManager.
     */
    public static void quotaRequestInProgress(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.quotaRequestStatus().getWithResponse("2B5C8515-37D8-4B6A-879B-CD641A2CF605",
            "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus",
            com.azure.core.util.Context.NONE);
    }
}
