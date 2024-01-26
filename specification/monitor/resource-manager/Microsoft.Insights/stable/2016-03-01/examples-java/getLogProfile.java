
import com.azure.core.util.Context;

/** Samples for LogProfiles Get. */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/getLogProfile.json
     */
    /**
     * Sample code: Get log profile.
     *
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void getLogProfile(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getLogProfiles().getWithResponse("default", Context.NONE);
    }
}
