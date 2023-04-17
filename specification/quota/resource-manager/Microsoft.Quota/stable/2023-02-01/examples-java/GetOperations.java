/** Samples for QuotaOperation List. */
public final class Main {
    /*
     * x-ms-original-file: specification/quota/resource-manager/Microsoft.Quota/stable/2023-02-01/examples/GetOperations.json
     */
    /**
     * Sample code: GetOperations.
     *
     * @param manager Entry point to QuotaManager.
     */
    public static void getOperations(com.azure.resourcemanager.quota.QuotaManager manager) {
        manager.quotaOperations().list(com.azure.core.util.Context.NONE);
    }
}
