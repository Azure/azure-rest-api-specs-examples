
/**
 * Samples for ManagedClusters GetGuardrailsVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/GetGuardrailsVersions.json
     */
    /**
     * Sample code: Get guardrails available versions.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        getGuardrailsAvailableVersions(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().getGuardrailsVersionsWithResponse("location1", "v1.0.0",
            com.azure.core.util.Context.NONE);
    }
}
