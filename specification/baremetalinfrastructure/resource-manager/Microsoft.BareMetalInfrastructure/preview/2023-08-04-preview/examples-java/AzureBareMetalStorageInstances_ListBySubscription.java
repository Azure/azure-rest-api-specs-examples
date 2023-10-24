/** Samples for AzureBareMetalStorageInstances List. */
public final class Main {
    /*
     * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalStorageInstances_ListBySubscription.json
     */
    /**
     * Sample code: List all AzureBareMetalStorage instances in a subscription.
     *
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void listAllAzureBareMetalStorageInstancesInASubscription(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager.azureBareMetalStorageInstances().list(com.azure.core.util.Context.NONE);
    }
}
