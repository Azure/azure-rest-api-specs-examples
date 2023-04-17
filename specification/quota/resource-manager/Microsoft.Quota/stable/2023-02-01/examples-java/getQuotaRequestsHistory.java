/** Samples for QuotaRequestStatus List. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/getQuotaRequestsHistory.json
     */
    /**
     * Sample code: QuotaRequestHistory.
     *
     * @param manager Entry point to QuotaManager.
     */
    public static void quotaRequestHistory(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager
            .quotaRequestStatus()
            .list(
                "subscriptions/D7EC67B3-7657-4966-BFFC-41EFD36BAAB3/providers/Microsoft.Compute/locations/eastus",
                null,
                null,
                null,
                com.azure.core.util.Context.NONE);
    }
}
