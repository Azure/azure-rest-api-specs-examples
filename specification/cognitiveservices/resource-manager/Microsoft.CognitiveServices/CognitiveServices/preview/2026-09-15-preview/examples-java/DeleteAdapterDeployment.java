
/**
 * Samples for AdapterDeployments Delete.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/DeleteAdapterDeployment.json
     */
    /**
     * Sample code: DeleteAdapterDeployment.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        deleteAdapterDeployment(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.adapterDeployments().delete("resourceGroupName", "accountName", "adapterDeploymentName",
            "\"0x8D1234567890ABC\"", com.azure.core.util.Context.NONE);
    }
}
