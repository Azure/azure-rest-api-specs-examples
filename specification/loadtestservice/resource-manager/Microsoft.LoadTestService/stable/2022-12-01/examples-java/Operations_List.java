
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/loadtestservice/resource-manager/Microsoft.LoadTestService/stable/2022-12-01/examples/
     * Operations_List.json
     */
    /**
     * Sample code: List the operations for the provider.
     * 
     * @param manager Entry point to LoadTestManager.
     */
    public static void listTheOperationsForTheProvider(com.azure.resourcemanager.loadtesting.LoadTestManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
