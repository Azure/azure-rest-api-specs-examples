import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/securityinsights/resource-manager/Microsoft.SecurityInsights/preview/2022-01-01-preview/examples/operations/ListOperations.json
     */
    /**
     * Sample code: Get all operations.
     *
     * @param manager Entry point to SecurityInsightsManager.
     */
    public static void getAllOperations(com.azure.resourcemanager.securityinsights.SecurityInsightsManager manager) {
        manager.operations().list(Context.NONE);
    }
}
