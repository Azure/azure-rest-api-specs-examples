/** Samples for AzureBareMetalInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalInstances_ListBySubscription.json
     */
    /**
     * Sample code: List all Azure Bare Metal Instances in a subscription.
     *
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void listAllAzureBareMetalInstancesInASubscription(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager.azureBareMetalInstances().list(com.azure.core.util.Context.NONE);
    }
}
