
/**
 * Samples for ManagedClusters ListGuardrailsVersions.
 */
public final class Main {
    /*
     * x-ms-original-file: 2026-04-02-preview/ListGuardrailsVersions.json
     */
    /**
     * Sample code: List Guardrails Versions.
     * 
     * @param manager Entry point to ContainerServiceManager.
     */
    public static void
        listGuardrailsVersions(com.azure.resourcemanager.containerservice.ContainerServiceManager manager) {
        manager.serviceClient().getManagedClusters().listGuardrailsVersions("location1",
            com.azure.core.util.Context.NONE);
    }
}
