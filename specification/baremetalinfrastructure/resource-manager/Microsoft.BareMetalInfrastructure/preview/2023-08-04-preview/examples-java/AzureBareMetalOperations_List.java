/** Samples for Operations List. */
public final class Main {
    /*
     * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalOperations_List.json
     */
    /**
     * Sample code: List all management operations supported by the AzureBareMetal RP.
     *
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void listAllManagementOperationsSupportedByTheAzureBareMetalRP(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager.operations().list(com.azure.core.util.Context.NONE);
    }
}
