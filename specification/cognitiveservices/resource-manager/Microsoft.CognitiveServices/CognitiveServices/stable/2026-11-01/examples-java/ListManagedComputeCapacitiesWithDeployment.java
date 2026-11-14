
/**
 * Samples for ManagedComputeCapacities List.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-11-01/ListManagedComputeCapacitiesWithDeployment.json
     */
    /**
     * Sample code: List Managed Compute Capacities For Deployment With Matching Account Location.
     * 
     * @param manager Entry point to CognitiveServicesManager.
     */
    public static void listManagedComputeCapacitiesForDeploymentWithMatchingAccountLocation(
        com.azure.resourcemanager.cognitiveservices.CognitiveServicesManager manager) {
        manager.managedComputeCapacities().list("MaaP", "H100_80GB",
            "/subscriptions/xxxxxxxx-xxxx-xxxx-xxxx-xxxxxxxxxxxx/resourceGroups/resourceGroupName/providers/Microsoft.CognitiveServices/accounts/accountName/managedComputeDeployments/deploymentName",
            "eastus", com.azure.core.util.Context.NONE);
    }
}
