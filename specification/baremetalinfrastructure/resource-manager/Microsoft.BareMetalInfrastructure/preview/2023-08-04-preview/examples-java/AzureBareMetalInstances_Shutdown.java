/** Samples for AzureBareMetalInstances Shutdown. */
public final class Main {
    /*
     * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalInstances_Shutdown.json
     */
    /**
     * Sample code: Shutdown an Azure Bare Metal Instance.
     *
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void shutdownAnAzureBareMetalInstance(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager
            .azureBareMetalInstances()
            .shutdown("myResourceGroup", "myABMInstance", com.azure.core.util.Context.NONE);
    }
}
