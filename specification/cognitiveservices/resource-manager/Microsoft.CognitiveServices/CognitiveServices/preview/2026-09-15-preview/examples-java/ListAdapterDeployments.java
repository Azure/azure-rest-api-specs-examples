
/**
 * Samples for AdapterDeployments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-09-15-preview/ListAdapterDeployments.json
     */
    /**
     * Sample code: ListAdapterDeployments.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listAdapterDeployments(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.adapterDeployments().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
