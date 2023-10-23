/** Samples for AzureBareMetalInstances GetByResourceGroup. */
public final class Main {
    /*
     * x-ms-original-file: specification/baremetalinfrastructure/resource-manager/Microsoft.BareMetalInfrastructure/preview/2023-08-04-preview/examples/AzureBareMetalInstances_Get.json
     */
    /**
     * Sample code: Get an Azure Bare Metal Instance.
     *
     * @param manager Entry point to BareMetalInfrastructureManager.
     */
    public static void getAnAzureBareMetalInstance(
        com.azure.resourcemanager.baremetalinfrastructure.BareMetalInfrastructureManager manager) {
        manager
            .azureBareMetalInstances()
            .getByResourceGroupWithResponse(
                "myResourceGroup", "myAzureBareMetalInstance", com.azure.core.util.Context.NONE);
    }
}
