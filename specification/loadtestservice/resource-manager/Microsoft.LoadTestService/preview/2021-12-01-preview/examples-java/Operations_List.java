import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/loadtestservice/resource-manager/Microsoft.LoadTestService/preview/2021-12-01-preview/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to LoadTestManager.
     */
    public static void operationsList(com.azure.resourcemanager.loadtestservice.LoadTestManager manager) {
        manager.operations().list(Context.NONE);
    }
}
