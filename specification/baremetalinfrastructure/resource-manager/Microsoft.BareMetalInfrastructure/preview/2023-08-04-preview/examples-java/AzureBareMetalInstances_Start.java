/** Samples for AzureBareMetalInstances Start. */
public final class Main {
    /*
     * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalInstances_Start.json
     */
    /**
     * Sample code: Start an Azure Bare Metal instance.
     *
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void startAnAzureBareMetalInstance(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager.azureBareMetalInstances().start("myResourceGroup", "myABMInstance", com.azure.core.util.Context.NONE);
    }
}
