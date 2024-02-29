
/**
 * Samples for LogProfiles List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/monitor/resource-manager/Microsoft.Insights/stable/2016-03-01/examples/listLogProfile.json
     */
    /**
     * Sample code: List log profiles.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void listLogProfiles(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.diagnosticSettings().manager().serviceClient().getLogProfiles().list(com.azure.core.util.Context.NONE);
    }
}
