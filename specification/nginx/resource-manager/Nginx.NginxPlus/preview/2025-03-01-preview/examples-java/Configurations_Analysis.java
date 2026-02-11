
/**
 * Samples for Configurations Analysis.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-03-01-preview/Configurations_Analysis.json
     */
    /**
     * Sample code: Configurations_Analysis.
     * 
     * @param manager Entry point to NginxManager.
     */
    public static void configurationsAnalysis(com.azure.resourcemanager.nginx.NginxManager manager) {
        manager.configurations().analysisWithResponse("myResourceGroup", "myDeployment", "default", null,
            com.azure.core.util.Context.NONE);
    }
}
