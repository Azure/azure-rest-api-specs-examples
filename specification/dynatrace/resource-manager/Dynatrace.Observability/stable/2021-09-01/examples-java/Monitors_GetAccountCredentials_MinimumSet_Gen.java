import com.azure.core.util.Context;

/** Samples for Monitors GetAccountCredentials. */
public final class Main {
    /*
     * x-ms-original-file: specification/dynatrace/resource-manager/Dynatrace.Observability/stable/2021-09-01/examples/Monitors_GetAccountCredentials_MinimumSet_Gen.json
     */
    /**
     * Sample code: Monitors_GetAccountCredentials_MinimumSet_Gen.
     *
     * @param manager Entry point to DynatraceManager.
     */
    public static void monitorsGetAccountCredentialsMinimumSetGen(
        com.azure.resourcemanager.dynatrace.DynatraceManager manager) {
        manager.monitors().getAccountCredentialsWithResponse("myResourceGroup", "myMonitor", Context.NONE);
    }
}
