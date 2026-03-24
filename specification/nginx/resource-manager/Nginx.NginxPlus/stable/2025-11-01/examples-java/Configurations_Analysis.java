
/**
 * Samples for Configurations Analysis.
 */
public final class Main {
    /*
     * x-ms-original-file: 2025-11-01/Configurations_Analysis.json
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
