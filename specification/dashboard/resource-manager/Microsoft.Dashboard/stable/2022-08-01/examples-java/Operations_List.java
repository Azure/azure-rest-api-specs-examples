import com.azure.core.util.Context;

/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/dashboard/resource-manager/Microsoft.Dashboard/stable/2022-08-01/examples/Operations_List.json
     */
    /**
     * Sample code: Operations_List.
     *
     * @param manager Entry point to DashboardManager.
     */
    public static void operationsList(com.azure.resourcemanager.dashboard.DashboardManager manager) {
        manager.operations().list(Context.NONE);
    }
}
