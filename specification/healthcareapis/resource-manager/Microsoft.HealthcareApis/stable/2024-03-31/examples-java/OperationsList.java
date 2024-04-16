
/**
 * Samples for Operations List.
 */
public final class Main {
    /*
     * x-ms-original-file:
     * specification/healthcareapis/resource-manager/Microsoft.HealthcareApis/stable/2024-03-31/examples/OperationsList.
     * json
     */
    /**
     * Sample code: List operations.
     * 
     * @param manager Entry point to HealthcareApisManager.
     */
    public static void listOperations(com.azure.resourcemanager.healthcareapis.HealthcareApisManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
