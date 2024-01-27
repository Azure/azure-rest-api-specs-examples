
/**
 * Samples for WebApps AnalyzeCustomHostname.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/web/resource-manager/Microsoft.Web/stable/2023-01-01/examples/AnalyzeCustomHostName.json
     */
    /**
     * Sample code: Analyze custom hostname for webapp.
     * 
     * @param azure The entry point for accessing resource management APIs in Azure.
     */
    public static void analyzeCustomHostnameForWebapp(com.azure.resourcemanager.AzureResourceManager azure) {
        azure.webApps().manager().serviceClient().getWebApps().analyzeCustomHostnameWithResponse("testrg123",
            "sitef6141", null, com.azure.core.util.Context.NONE);
    }
}
