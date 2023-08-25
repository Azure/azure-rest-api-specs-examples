import com.azure.resourcemanager.apimanagement.models.QuotaCounterValueUpdateContract;

/** Samples for QuotaByCounterKeys Update. */
public final class Main {
    /*
     * x-ms-original-file: specification/apimanagement/resource-manager/Microsoft.ApiManagement/stable/2022-08-01/examples/ApiManagementUpdateQuotaCounterKey.json
     */
    /**
     * Sample code: ApiManagementUpdateQuotaCounterKey.
     *
     * @param manager Entry point to ApiManagementManager.
     */
    public static void apiManagementUpdateQuotaCounterKey(
        com.azure.resourcemanager.apimanagement.ApiManagementManager manager) {
        manager
            .quotaByCounterKeys()
            .updateWithResponse(
                "rg1",
                "apimService1",
                "ba",
                new QuotaCounterValueUpdateContract().withCallsCount(0).withKbTransferred(2.5630078125D),
                com.azure.core.util.Context.NONE);
    }
}
