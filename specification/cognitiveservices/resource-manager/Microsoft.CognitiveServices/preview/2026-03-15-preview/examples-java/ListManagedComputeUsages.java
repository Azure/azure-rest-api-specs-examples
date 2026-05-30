
/**
 * Samples for ManagedComputeUsagesOperationGroup List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-03-15-preview/ListManagedComputeUsages.json
     */
    /**
     * Sample code: List Managed Compute Usages.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void
        listManagedComputeUsages(com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedComputeUsagesOperationGroups().list("eastus", com.azure.core.util.Context.NONE);
    }
}
