
/**
 * Samples for ManagedComputeDeployments List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListVmManagedComputeDeployments.json
     */
    /**
     * Sample code: ListVmManagedComputeDeployments.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listVmManagedComputeDeployments(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedComputeDeployments().list("resourceGroupName", "accountName", com.azure.core.util.Context.NONE);
    }
}
