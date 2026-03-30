
/**
 * Samples for WebApps AnalyzeCustomHostname.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-05-01/AnalyzeCustomHostName.json
     */
    /**
     * Sample code: Analyze custom hostname for webapp.
     * 
     * @param manager Entry point to AppServiceManager.
     */
    public static void analyzeCustomHostnameForWebapp(com.azure.resourcemanager.appservice.AppServiceManager manager) {
        manager.serviceClient().getWebApps().analyzeCustomHostnameWithResponse("testrg123", "sitef6141", null,
            com.azure.core.util.Context.NONE);
    }
}
