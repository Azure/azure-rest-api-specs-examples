import com.azure.core.util.Context;

/** Samples for ProviderInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/workloads/resource-manager/Microsoft.Workloads/preview/2021-12-01-preview/examples/workloadmonitor/ProviderInstances_List.json
     */
    /**
     * Sample code: List all SAP monitors providers in a subscription.
     *
     * @param manager Entry point to WorkloadsManager.
     */
    public static void listAllSAPMonitorsProvidersInASubscription(
        com.azure.resourcemanager.workloads.WorkloadsManager manager) {
        manager.providerInstances().list("myResourceGroup", "mySapMonitor", Context.NONE);
    }
}
