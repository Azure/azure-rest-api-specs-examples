
import com.azure.core.util.Context;

/** Samples for VMInsights GetOnboardingStatus. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/preview/2018-11-27-preview/examples/
     * getOnboardingStatusSingleVMUnknown.json
     */
    /**
     * Sample code: Get status for a VM that has not yet reported data.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getStatusForAVMThatHasNotYetReportedData(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getVMInsights().getOnboardingStatusWithResponse(
            "subscriptions/3d51de47-8d1c-4d24-b42f-bcae075dfa87/resourceGroups/vm-resource-group/providers/Microsoft.Compute/virtualMachines/ubuntu-vm",
            Context.NONE);
    }
}
