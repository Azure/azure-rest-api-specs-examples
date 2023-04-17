/** Samples for QuotaRequestStatus Get. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getQuotaRequestStatusFailed.json
     */
    /**
     * Sample code: QuotaRequestFailed.
     *
     * @param manager Entry point to QuotaManager.
     */
    public static void quotaRequestFailed(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager
            .quotaRequestStatus()
            .getWithResponse(
                "2B5C8515-37D8-4B6A-879B-CD641A2CF605",
                "subscriptions/00000000-0000-0000-0000-000000000000/providers/Microsoft.Compute/locations/eastus",
                com.azure.core.util.Context.NONE);
    }
}
